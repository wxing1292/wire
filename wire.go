//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/wxing1292/wire/common/cache"
	"github.com/wxing1292/wire/common/logger"
	"github.com/wxing1292/wire/common/metrics"
	"github.com/wxing1292/wire/common/service"
)

func WireService() (*service.ServiceImpl, error) {
	wire.Build(
		logger.WireSet,
		metrics.WireSet,
		cache.WireSet,
		service.WireSet,
	)
	return &service.ServiceImpl{}, nil
}
