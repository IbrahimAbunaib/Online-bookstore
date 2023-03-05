package router

import (
	"github.com/IbrahimAbunaib/go-mux-api/Documents/Clean_Architect_Api/interface/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Get("/users",
		func(context echo.Context) error { return c.User.GetUsers(context) })

	return e
}
