package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://wishy.aandrku.dev"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

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
