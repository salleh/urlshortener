package handler

import (
	"context"
	"net/http"
	"urlshortener/component"
	"urlshortener/ent"
	"urlshortener/ent/url"
	utils "urlshortener/utility"

	"github.com/labstack/echo/v4"
)

func RedirectById(c echo.Context) error {
	var err error
	var entClient *ent.Client
	entClient, err = component.GetEntClient()
	if err != nil {
		utils.AppLogger.Errorf("Entity Client Fetching Error: %v", err.Error())

		return c.String(http.StatusInternalServerError, "Sorry, something wrong here. Please try again.")
	}

	// ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.GetUnifiedConfig().AppTimeout)*time.Second)
	// defer cancel()
	ctx := context.Background()
	var result *ent.Url
	result, err = entClient.Url.Query().
		Where(url.ShortPathEQ(c.Param("id"))).
		Only(ctx)
	if err != nil {
		utils.AppLogger.Errorf("Entity Operation Error: %v", err.Error())
		switch err.(type) {
		case *ent.NotFoundError:
			return c.String(http.StatusNotFound, "Could not find the url requested")
		default:
			return c.String(http.StatusInternalServerError, "Something wrong happened. Please try later")
		}
	}

	return c.Redirect(http.StatusMovedPermanently, result.LongPath)
}
