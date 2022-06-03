package cache

import (
	"context"
	"math/bits"
	"time"
)

// Option is a function which initializes config.
type Option func(conf *config)

// applyOptions applies opts to conf.
// 应用选项 用在申请cache
func applyOptions(conf *config, opts []Option) *config {
	for _, opt := range opts {
		opt(conf)
	}
	return conf
}

// WithMapSize is an option setting initializing map size of cache.
func WithMapSize(mapSize uint) Option {
	return func(conf *config) {
		conf.mapSize = int(mapSize)
	}
}

// WithSegmentSize is an option setting initializing segment size of cache.
// segmentSize must be the pow of 2 (such as 64) or the segments may be uneven.
func WithSegmentSize(segmentSize uint) Option {
	if bits.OnesCount(segmentSize) > 1 {
		panic("cachego: segmentSize must be the pow of 2 (such as 64) or the segments may be uneven.")
	}

	return func(conf *config) {
		conf.segmentSize = int(segmentSize)
	}
}

// WithAutoGC is an option turning on automatically gc.
//WithAutoGC 是一个自动开启 gc 的选项
func WithAutoGC(d time.Duration) Option {
	return func(conf *config) {
		if d > 0 {
			conf.gcDuration = d
		}
	}
}

// WithAutoGC is an option turning on automatically gc.
//WithAutoGC 是一个自动开启 gc 的选项
func WithAutoRandGC(d time.Duration) Option {
	return func(conf *config) {
		if d > 0 {
			conf.gcRandDuration = d
		}
	}
}

// WithDisableSingleflight is an option disabling single-flight mode of cache.
func WithDisableSingleflight() Option {
	return func(conf *config) {
		conf.singleflight = false
	}
}

// OpOption is a function which initializes opConfig.
type OpOption func(conf *opConfig)

// applyOpOptions applies opts to conf.
func applyOpOptions(conf *opConfig, opts []OpOption) *opConfig {
	for _, opt := range opts {
		opt(conf)
	}
	return conf
}

// WithOpContext sets context to ctx.
func WithOpContext(ctx context.Context) OpOption {
	return func(conf *opConfig) {
		conf.ctx = ctx
	}
}

// WithOpTTL sets the ttl of missed key if loaded to ttl.
func WithOpTTL(ttl time.Duration) OpOption {
	return func(conf *opConfig) {
		conf.ttl = ttl
	}
}

// WithOpNoTTL sets the ttl of missed key to no ttl.
func WithOpNoTTL() OpOption {
	return func(conf *opConfig) {
		conf.ttl = noTTL
	}
}

// WithOpOnMissed sets onMissed to Get operation.
func WithOpOnMissed(onMissed func(ctx context.Context, key interface{}) (data interface{}, err error)) OpOption {
	return func(conf *opConfig) {
		conf.onMissed = onMissed
	}
}

// WithOpDisableSingleflight sets the single-flight mode to false.
func WithOpDisableSingleflight() OpOption {
	return func(conf *opConfig) {
		conf.singleflight = false
	}
}

// WithOpDisableReload sets the reloading flag to false.
func WithOpDisableReload() OpOption {
	return func(conf *opConfig) {
		conf.reload = false
	}
}
