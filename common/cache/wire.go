//go:build wireinject

package cache

import (
	"github.com/google/wire"
	"github.com/wxing1292/wire/common/logger"
	"github.com/wxing1292/wire/common/metrics"
)

var WireSet = wire.NewSet(
	wire.Bind(new(Cache), new(*CacheImpl)),
	StartCache,
)

func StartCache(
	logger logger.Logger,
	metrics metrics.Metrics,
) (*CacheImpl, error) {
	cache, err := wireCache(logger, metrics)
	if err != nil {
		return nil, err
	}

	if err := cache.Start(); err != nil {
		return nil, cache.Stop()
	}
	return cache, nil
}

// TODO: how do we stop it?
func StopCache() {}

func wireCache(
	logger logger.Logger,
	metrics metrics.Metrics,
) (*CacheImpl, error) {
	wire.Build(NewCache)
	return &CacheImpl{}, nil
}
