//go:build wireinject

package service

// TODO: there is no built in compiler support to auto import
//  due to tag `//go:build wireinject`
import (
	"github.com/google/wire"
	"github.com/wxing1292/wire/common/cache"
	"github.com/wxing1292/wire/common/logger"
	"github.com/wxing1292/wire/common/metrics"
)

var WireSet = wire.NewSet(
	StartService,
)

func StartService(
	cache cache.Cache,
	logger logger.Logger,
	metrics metrics.Metrics,
) (*ServiceImpl, error) {
	service, err := wireService(cache, logger, metrics)
	if err != nil {
		return nil, err
	}

	if err := service.Start(); err != nil {
		return nil, service.Stop()
	}
	return service, nil
}

// TODO: there is no built in compiler support, e.g.
//  try `func stopService()` instead of `func stopService() {}`
//  due to tag `//go:build wireinject`
// TODO: how do we stop it?
func StopService() {}

func wireService(
	cache cache.Cache,
	logger logger.Logger,
	metrics metrics.Metrics,
) (*ServiceImpl, error) {
	wire.Build(
		NewService,
	)
	return &ServiceImpl{}, nil
}
