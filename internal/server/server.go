package server

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ugursogukpinar/go-api-boilerplate/internal/config"
	"github.com/ugursogukpinar/go-api-boilerplate/internal/server/controllers"
	"go.uber.org/fx"
)

func NewServer(logger *slog.Logger) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Gzip())

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	e.Use(middleware.Recover())

	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "",
		ContentTypeNosniff:    "",
		XFrameOptions:         "",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "default-src 'self'",
	}))

	return e
}

func RegisterHooks(lc fx.Lifecycle, logger *slog.Logger, cfg *config.Config, e *echo.Echo) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("Starting server.", slog.String("addr", cfg.HTTP.ListenAddress))
			go e.Start(cfg.HTTP.ListenAddress)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			e.Shutdown(ctx)
			return nil
		},
	})
}

var Module = fx.Options(fx.Provide(NewServer), controllers.Module, fx.Invoke(RegisterHooks))
