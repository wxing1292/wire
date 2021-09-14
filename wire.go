//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/wxing1292/wire/common/logger"
	"github.com/wxing1292/wire/common/metrics"
	"github.com/wxing1292/wire/common/service"
)

func WireService() (*service.ServiceImpl, error) {
	wire.Build(
		logger.WireSet,
		metrics.WireSet,
		service.WireService,
	)
	return &service.ServiceImpl{}, nil
}
