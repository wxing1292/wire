//go:build wireinject

package service

import (
	"github.com/google/wire"
	"github.com/wxing1292/wire/common/logger"
	"github.com/wxing1292/wire/common/metrics"
)

func WireService(
	logger logger.Logger,
	metrics metrics.Metrics,
) (*ServiceImpl, error) {
	wire.Build(
		NewService,
	)
	return &ServiceImpl{}, nil
}
