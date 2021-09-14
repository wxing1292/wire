//go:build wireinject

package logger

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	wire.Bind(new(Logger), new(*LoggerImpl)),
	WireLogger,
)

func WireLogger() (*LoggerImpl, error) {
	wire.Build(NewLogger)
	return &LoggerImpl{}, nil
}
