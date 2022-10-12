package main

import (
	"context"
	"fmt"
	"net/http"

	"urlshortener/appconst"
	"urlshortener/component"
	"urlshortener/config"
	"urlshortener/router"
	utils "urlshortener/utility"
)

func init() {
	fmt.Printf("Initializing URL Shortener Server version %s\n", appconst.AppVersion)
}
func main() {
	// Change logging level of the app according to configuration
	utils.ChangeLogLevel(config.GetUnifiedConfig().AppLogLevel)

	// Change random generator library meta value
	utils.SetIDMeta(config.GetUnifiedConfig().AppIDMeta)

	ec, err := component.GetEntClient()
	if err != nil {
		utils.AppLogger.Fatal("Ent Client fetching failed. Terminating")
	}

	// Run the auto migration tool.
	utils.AppLogger.Info("Running schema migration on entity repository")
	if err = ec.Schema.Create(context.Background()); err != nil {
		utils.AppLogger.Errorf("Failed creating schema resources: %v", err)
		utils.AppLogger.Warn("Schema migration failed. Data integrity of the service may be at risk!!!")
	}
	utils.AppLogger.Info("Schema migration/update on entity repository completed")

	e, err := component.GetEchoInstance()
	if err != nil {
		utils.AppLogger.Fatal("Echo Server instance fetching failed. Terminating")
	}

	// Register route
	utils.AppLogger.Info("Registering route for Echo Server instance...")
	utils.AppLogger.Info("Registering universal route")
	err = router.RegisterUniversalRoute(e)
	if err != nil {
		utils.AppLogger.Fatal("Universal route registration failed. Terminating")
	}
	utils.AppLogger.Info("Universal route registration completed")

	utils.AppLogger.Info("Echo Server instance setup completed")

	// Start Echo Server
	utils.AppLogger.Info("Starting up URL Shortener HTTP Server")

	if err = e.Start(fmt.Sprintf("%s:%d", config.GetUnifiedConfig().HTTPListen, config.GetUnifiedConfig().HTTPPort)); err != http.ErrServerClosed {
		utils.AppLogger.Fatalf("Service terminated: %v", err)
	}

	utils.AppLogger.Info("OTP Service Ended")

}
