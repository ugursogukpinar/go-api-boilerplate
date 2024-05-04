package main

import (
	"log/slog"

	"github.com/ugursogukpinar/go-api-boilerplate/internal/config"
	"github.com/ugursogukpinar/go-api-boilerplate/internal/logger"
	"github.com/ugursogukpinar/go-api-boilerplate/internal/server"
	"github.com/ugursogukpinar/go-api-boilerplate/internal/services"
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
		services.Module,
		server.Module,
	)

	app.Run()
}
