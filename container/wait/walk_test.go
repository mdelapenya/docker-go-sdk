package wait_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/docker/go-sdk/container/wait"
)

type testDefinition struct {
	WaitingFor wait.Strategy
}

func TestWalk(t *testing.T) {
	def := testDefinition{
		WaitingFor: wait.ForAll(
			wait.ForFile("/tmp/file"),
			wait.ForHTTP("/health"),
			wait.ForAll(
				wait.ForFile("/tmp/other"),
			),
		),
	}

	t.Run("walk", func(t *testing.T) {
		var count int
		err := wait.Walk(&def.WaitingFor, func(_ wait.Strategy) error {
			count++
			return nil
		})
		require.NoError(t, err)
		require.Equal(t, 5, count)
	})

	t.Run("stop", func(t *testing.T) {
		var count int
		err := wait.Walk(&def.WaitingFor, func(_ wait.Strategy) error {
			count++
			return wait.ErrVisitStop
		})
		require.NoError(t, err)
		require.Equal(t, 1, count)
	})

	t.Run("remove", func(t *testing.T) {
		// walkRemoveFileStrategy {
		var count, matched int
		err := wait.Walk(&def.WaitingFor, func(s wait.Strategy) error {
			count++
			if _, ok := s.(*wait.FileStrategy); ok {
				matched++
				return wait.ErrVisitRemove
			}

			return nil
		})
		// }
		require.NoError(t, err)
		require.Equal(t, 5, count)
		require.Equal(t, 2, matched)

		count = 0
		matched = 0
		err = wait.Walk(&def.WaitingFor, func(s wait.Strategy) error {
			count++
			if _, ok := s.(*wait.FileStrategy); ok {
				matched++
			}
			return nil
		})
		require.NoError(t, err)
		require.Equal(t, 3, count)
		require.Zero(t, matched)
	})

	t.Run("remove-stop", func(t *testing.T) {
		def := testDefinition{
			WaitingFor: wait.ForAll(
				wait.ForFile("/tmp/file"),
				wait.ForHTTP("/health"),
			),
		}
		var count int
		err := wait.Walk(&def.WaitingFor, func(_ wait.Strategy) error {
			count++
			return errors.Join(wait.ErrVisitRemove, wait.ErrVisitStop)
		})
		require.NoError(t, err)
		require.Equal(t, 1, count)
		require.Nil(t, def.WaitingFor)
	})

	t.Run("nil-root", func(t *testing.T) {
		err := wait.Walk(nil, func(_ wait.Strategy) error {
			return nil
		})
		require.EqualError(t, err, "root strategy is nil")
	})

	t.Run("direct-single", func(t *testing.T) {
		def := testDefinition{
			WaitingFor: wait.ForFile("/tmp/file"),
		}
		requireVisits(t, def, 1)
	})

	t.Run("for-all-single", func(t *testing.T) {
		def := testDefinition{
			WaitingFor: wait.ForAll(
				wait.ForFile("/tmp/file"),
			),
		}
		requireVisits(t, def, 2)
	})
}

// requireVisits validates the number of visits for a given request.
func requireVisits(t *testing.T, def testDefinition, expected int) {
	t.Helper()

	var count int
	err := wait.Walk(&def.WaitingFor, func(_ wait.Strategy) error {
		count++
		return nil
	})
	require.NoError(t, err)
	require.Equal(t, expected, count)
}
