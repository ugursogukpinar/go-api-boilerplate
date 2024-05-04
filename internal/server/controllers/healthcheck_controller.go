package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/ugursogukpinar/go-api-boilerplate/internal/services"
)

type HealtcheckController struct {
	myService *services.MyService
}

func NewHealthcheckController(s *services.MyService) HealtcheckController {
	return HealtcheckController{
		myService: s,
	}
}

func (cnt HealtcheckController) Get(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"status":  true,
		"message": cnt.myService.Greeting("World!"),
	})
}
