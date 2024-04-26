package main

import (
	"log/slog"

	"github.com/ugursogukpinar/go-api-boilerplate/modules/config"
	"github.com/ugursogukpinar/go-api-boilerplate/modules/logger"
	"github.com/ugursogukpinar/go-api-boilerplate/modules/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	app := fx.New(
		logger.Module,
		fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
			return &fxevent.SlogLogger{Logger: logger}
		}),
		config.Module,
		server.Module,
	)

	app.Run()
}
