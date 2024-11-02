package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "main Go app")
	})
	app.Logger.Fatal(app.Start(":1011"))
}
