package dockercontainer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFile_validate(t *testing.T) {
	t.Run("empty-reader", func(t *testing.T) {
		file := File{}
		err := file.validate()
		require.Error(t, err)
	})

	t.Run("empty-container-path", func(t *testing.T) {
		file := File{
			Reader: bytes.NewReader([]byte("test")),
		}
		err := file.validate()
		require.Error(t, err)
	})

	t.Run("valid", func(t *testing.T) {
		file := File{
			Reader:        bytes.NewReader([]byte("test")),
			ContainerPath: "test",
		}
		err := file.validate()
		require.NoError(t, err)
	})
}
