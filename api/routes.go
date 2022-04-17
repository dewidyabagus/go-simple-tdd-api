package app

import (
	"github.com/labstack/echo/v4"

	"go-simple-api/api/v1/welcome"
)

type Routes struct {
	Welcome *welcome.Controller
}

func CreateRoutes(e *echo.Echo, routes *Routes) {
	v1 := e.Group("/v1")

	v1.GET("/welcome", routes.Welcome.GetIndex)
}
