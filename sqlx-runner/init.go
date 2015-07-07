package runner

import (
	"time"

	"github.com/mgutz/logxi/v1"
	"github.com/JustinTulloss/dat"
	"github.com/JustinTulloss/dat/kvs"
	"github.com/JustinTulloss/dat/postgres"
)

var logger log.Logger

// LogQueriesThreshold is the threshold for logging "slow" queries
var LogQueriesThreshold time.Duration

func init() {
	dat.Dialect = postgres.New()
	logger = log.New("dat:sqlx")
}

// Cache caches query results.
var Cache kvs.KeyValueStore

// SetCache sets this runner's cache. The default cache is in-memory
// based. See cache.MemoryKeyValueStore.
func SetCache(store kvs.KeyValueStore) {
	Cache = store
}
