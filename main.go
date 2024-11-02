package main

import (
	"go-mail-sync/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.Logger())
	app.GET("/", func(c echo.Context) error {
		inboxs, _ := services.GetInbox()
		return c.JSON(http.StatusOK, inboxs)
	})
	app.Logger.Fatal(app.Start(":1011"))
}
