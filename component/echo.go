package component

import (
	"strings"

	"urlshortener/config"
	utils "urlshortener/utility"

	"github.com/brpaz/echozap"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	echoInstance *echo.Echo
)

func initEchoInstance() (*echo.Echo, error) {
	utils.AppLogger.Info("Initializing Echo Server instance...")
	e, err := NewEchoInstance()
	if err != nil {
		utils.AppLogger.Errorf("Unable to initialize Echo Server Instance: %v", err)
		return nil, err
	}
	echoInstance = e
	utils.AppLogger.Info("Echo Server instance initialized")

	return e, nil
}

func GetEchoInstance() (*echo.Echo, error) {
	utils.AppLogger.Debug("Fetching Echo Server instance...")
	var err error
	if echoInstance == nil {
		utils.AppLogger.Debug("Echo Server instance does not exist. Attempting to initialize a new instance...")
		echoInstance, err = initEchoInstance()
		if err != nil {
			utils.AppLogger.Errorf("Fail to initialize Echo Server Instance: %v", err)
			return nil, err
		}
		utils.AppLogger.Debug("Echo Server instance Initialized")
	}

	return echoInstance, err
}

func NewEchoInstance() (*echo.Echo, error) {
	utils.AppLogger.Info("Creating new Echo Server instance...")
	e := echo.New()

	if config.GetUnifiedConfig().AppLogLevel == "debug" {
		e.Debug = true
	}

	e.HideBanner = !config.GetUnifiedConfig().AppShowBanner

	e.Use(middleware.Recover())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return strings.Replace(uuid.New().String(), "-", "", -1)
		},
	}))
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            3600,
		HSTSExcludeSubdomains: false,
		ContentSecurityPolicy: "default-src 'self'",
		CSPReportOnly:         false,
		ReferrerPolicy:        "origin-when-cross-origin",
	}))

	p := GetPrometheusInstance()
	p.Use(e)

	e.Use(echozap.ZapLogger(utils.Logger))
	utils.AppLogger.Info("New Echo Server instance created")

	return e, nil
}
