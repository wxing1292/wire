//go:build wireinject

package metrics

import (
	"github.com/google/wire"
	"github.com/wxing1292/wire/common/logger"
)

var WireSet = wire.NewSet(
	wire.Bind(new(Metrics), new(*MetricsImpl)),
	WireMetrics,
)

func WireMetrics(logger logger.Logger) (*MetricsImpl, error) {
	wire.Build(NewMetrics)
	return &MetricsImpl{}, nil
}
