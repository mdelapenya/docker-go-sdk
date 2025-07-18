package config

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_AuthConfigsForImages(t *testing.T) {
	config := Config{
		AuthConfigs: map[string]AuthConfig{
			"registry1.io": {Username: "user1", Password: "pass1"},
			"registry2.io": {Username: "user2", Password: "pass2"},
		},
	}

	images := []string{
		"registry1.io/repo/image:tag",
		"registry2.io/repo/image:tag",
	}

	authConfigs, err := config.AuthConfigsForImages(images)
	require.NoError(t, err)
	require.Len(t, authConfigs, 2)
	require.Equal(t, "user1", authConfigs["registry1.io"].Username)
	require.Equal(t, "user2", authConfigs["registry2.io"].Username)

	// Verify caching worked
	stats := config.cacheStats()
	require.Equal(t, 2, stats.Size)
}

func TestConfig_CacheManagement(t *testing.T) {
	config := Config{
		AuthConfigs: map[string]AuthConfig{
			"test.io": {Username: "user", Password: "pass"},
		},
	}

	t.Run("cache-initialization", func(t *testing.T) {
		stats := config.cacheStats()
		require.Equal(t, 0, stats.Size)
		require.NotEmpty(t, stats.CacheKey)
	})

	t.Run("cache-population", func(t *testing.T) {
		_, err := config.AuthConfigForHostname("test.io")
		require.NoError(t, err)

		stats := config.cacheStats()
		require.Equal(t, 1, stats.Size)
	})

	t.Run("cache-clearing", func(t *testing.T) {
		config.clearAuthCache()
		stats := config.cacheStats()
		require.Equal(t, 0, stats.Size)
	})
}

func TestConfig_ConcurrentAccess(t *testing.T) {
	config := Config{
		AuthConfigs: map[string]AuthConfig{
			"test.io": {Username: "user", Password: "pass"},
		},
	}

	const numGoroutines = 10
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			_, err := config.AuthConfigForHostname("test.io")
			require.NoError(t, err)
		}()
	}

	wg.Wait()
	stats := config.cacheStats()
	require.Equal(t, 1, stats.Size)
}

func TestConfig_CacheKeyGeneration(t *testing.T) {
	config1 := Config{
		AuthConfigs: map[string]AuthConfig{
			"test.io": {Username: "user1", Password: "pass1"},
		},
	}

	config2 := Config{
		AuthConfigs: map[string]AuthConfig{
			"test.io": {Username: "user2", Password: "pass2"},
		},
	}

	stats1 := config1.cacheStats()
	stats2 := config2.cacheStats()

	require.NotEqual(t, stats1.CacheKey, stats2.CacheKey)
}
