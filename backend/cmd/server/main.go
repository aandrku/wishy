package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		m := struct {
			Message string `json:"message"`
		}{
			Message: "Hello, World",
		}

		return c.JSON(http.StatusOK, m)
	})
	e.Logger.Fatal(e.Start("0.0.0.0:3000"))
}
