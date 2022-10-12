package component

import (
	"urlshortener/config"
	utils "urlshortener/utility"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
)

var (
	prometheusInstance *prometheus.Prometheus
)

// urlSkipper middleware ignores metrics on some route
func PrometheusUrlSkipper(c echo.Context) bool {

	// if strings.HasPrefix(c.Path(), "/") {
	// 	return true
	// }

	// if strings.HasPrefix(c.Path(), "/version") {
	// 	return true
	// }

	return config.PathToSkip[c.Path()]
}

func GetPrometheusInstance() *prometheus.Prometheus {
	if prometheusInstance == nil {
		utils.AppLogger.Info("Creating new prometheus middleware instance")
		prometheusInstance = prometheus.NewPrometheus("otp", PrometheusUrlSkipper)
	}

	return prometheusInstance
}
