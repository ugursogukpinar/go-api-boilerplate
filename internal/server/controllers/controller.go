package controllers

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Controller interface{}

func AsController(f any) any {
	return fx.Annotate(f, fx.As(new(Controller)), fx.ResultTags(`group:"controllers"`))
}

func BindRoutes(controllers []Controller, e *echo.Echo, logger *slog.Logger) {
	logger.Info("Binding controllers", slog.Int("count", len(controllers)))

	for _, c := range controllers {
		switch cnt := c.(type) {
		case HealtcheckController:
			e.GET("/healthcheck", cnt.Get)
		}
	}
}

var Module = fx.Options(
	fx.Provide(
		AsController(NewHealthcheckController),
	),
	fx.Invoke(fx.Annotate(BindRoutes, fx.ParamTags(`group:"controllers"`))),
)
