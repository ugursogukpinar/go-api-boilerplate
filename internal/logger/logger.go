package logger

import (
	"log/slog"
	"os"

	"go.uber.org/fx"
)

func NewLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return logger
}

var Module = fx.Options(fx.Provide(NewLogger))
