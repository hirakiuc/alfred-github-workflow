package cache

import "time"

const (
	maxCacheAgeDays = 7 // How long to cache data
)

func getMaxCacheAge() time.Duration {
	return time.Duration(maxCacheAgeDays) * (time.Second * 60 * 60 * 24)
}
