package main

import (
	"context"

	"github.com/ajayjadhav201/golang-templ/common"
	"github.com/ajayjadhav201/golang-templ/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	//
	common.Println("App started")
	app := echo.New()
	//
	app.Use(WithUser)
	//
	routes.RegisterRoutes(app)
	//
	app.Start(":80")

}

func WithUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ct := context.WithValue(c.Request().Context(), "user", "ajay@gmail.com")
		c.SetRequest(c.Request().WithContext(ct))
		return next(c)
	}
}
