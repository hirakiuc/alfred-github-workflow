package cache

import "time"

const (
	maxCacheAgeMinutes = 1 // How long to cache data
)

func getMaxCacheAge() time.Duration {
	return maxCacheAgeMinutes * time.Minute
}
