package cache

import (
	"context"
	"time"
)

const (
	noTTL = 0
)

// cache的config
type config struct {
	mapSize int

	segmentSize int

	gcDuration time.Duration

	gcRandDuration time.Duration

	singleflight bool
}

// newDefaultConfig returns the default config of cache.
func newDefaultConfig() *config {
	return &config{
		mapSize:        128,
		segmentSize:    128,
		gcDuration:     0,
		gcRandDuration: 0,
		singleflight:   true,
	}
}

// 每一个操作算子的config
type opConfig struct {
	// ctx is the context of operation.
	ctx context.Context

	// ttl is the ttl of entry set to the cache in operation.
	ttl time.Duration

	// onMissed is the function which will be called if not nil in operation.
	onMissed func(ctx context.Context, key interface{}) (data interface{}, err error)

	// singleflight means the call of onMissed is single-flight mode.
	// This is a recommended way to load data from storages to cache, however,
	// it may decrease the success rate of loading data.
	singleflight bool

	// reload means this operation will reload data from onMissed to cache.
	reload bool
}

// newDefaultGetConfig returns the default config of Get operations.
func newDefaultGetConfig() *opConfig {
	return &opConfig{
		ctx:          context.Background(),
		ttl:          10 * time.Second,
		onMissed:     nil,
		singleflight: true,
		reload:       true,
	}
}

// newDefaultSetConfig returns the default config of Set operations.
func newDefaultSetConfig() *opConfig {
	return &opConfig{
		ttl: noTTL,
	}
}
