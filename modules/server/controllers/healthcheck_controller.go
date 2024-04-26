package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealtcheckController struct{}

func NewHealthcheckController() HealtcheckController {
	return HealtcheckController{}
}

func (cnt HealtcheckController) Get(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"status": true,
	})
}
