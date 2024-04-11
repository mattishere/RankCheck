package globals

import (
	"RankCheck/cache"
	"RankCheck/notifs"
	"os"
	"strconv"
	"time"
)

var (
	Cooldowns  = cache.NewCache(10)   // 10 seconds
	StatsCache = cache.NewCache(3600) // 1 hour
)

func UpdateCaches() {
	updateString := os.Getenv("CACHE_UPDATE")

	CacheUpdateInterval, err := strconv.Atoi(updateString)
	if err != nil {
		notifs.Error("Error parsing CACHE_UPDATE environment variable: " + err.Error())
	}
	for {
		time.Sleep(time.Duration(CacheUpdateInterval) * time.Second)
		Cooldowns.Update()
		StatsCache.Update()
		notifs.Background("Caches updated")
	}
}
