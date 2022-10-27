package router

import (
	"net/http"
	"urlshortener/appconst"
	"urlshortener/handler"

	"github.com/labstack/echo/v4"
)

func RegisterUniversalRoute(e *echo.Echo) error {
	var err error

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Assalamualaikum Dunia dari URL Shortener Service")
	})

	e.GET("/:id", handler.RedirectById)

	e.GET("/version", func(c echo.Context) error {
		return c.String(http.StatusOK, appconst.AppVersion)
	})

	return err
}
