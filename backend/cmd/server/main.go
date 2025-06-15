package main

import (
	"flag"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	tls := flag.Bool("tls", false, "Usage: -tls")
	port := flag.String("port", "3000", "Usage: -port=3000")
	flag.Parse()

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

	if *tls {
		e.Logger.Fatal(e.StartTLS(":"+*port, "tls/certFile", "tls/keyFile"))
	} else {
		e.Logger.Fatal(e.Start(":" + *port))
	}

}
