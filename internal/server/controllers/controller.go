package controllers

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ugursogukpinar/go-api-boilerplate/internal/services"
	"go.uber.org/fx"
)

type Controller interface{}

func AsController(f any) any {
	return fx.Annotate(f, fx.As(new(Controller)), fx.ResultTags(`group:"controllers"`))
}

func BindRoutes(controllers []Controller, e *echo.Echo, logger *slog.Logger, authService *services.AuthService) {
	logger.Info("Binding controllers", slog.Int("count", len(controllers)))

	authGroup := e.Group("/auth")
	authGroup.Any("*", echo.WrapHandler(authService.GetAuthRoutes()))
	e.GET("/auth/deneme", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status": true,
		})
	})

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
