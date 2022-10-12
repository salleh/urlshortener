package component

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"urlshortener/config"
	"urlshortener/ent"
	utils "urlshortener/utility"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	entClient *ent.Client
)

func initEntClient() (*ent.Client, error) {
	utils.AppLogger.Infof("Initializing Ent Client instance...")
	ec, err := NewEntClient()
	if err != nil {
		utils.AppLogger.Errorf("Unable to initialize Ent Client: %v", err)
		return nil, err
	}
	entClient = ec
	utils.AppLogger.Info("Ent Client initialized")

	return ec, nil
}

func GetEntClient() (*ent.Client, error) {
	utils.AppLogger.Debug("Fetching Repository Client")
	var err error
	if entClient == nil {
		utils.AppLogger.Debug("No Entity Repository Client found. Creating new instance.")
		entClient, err = initEntClient()
		if err != nil {
			utils.AppLogger.Errorf("Unable to create Ent Client: %v", err)
			return nil, err
		}
		utils.AppLogger.Debug("Entity Repository Client Initialized")

	}

	return entClient, err
}

func getDBConnStr(config *config.UnifiedConfig) string {
	ex, err := os.Executable()
	if err != nil {
		utils.AppLogger.Fatalf("Unable to determine current path: %v", err)
	}
	exPath := filepath.Dir(ex)
	return fmt.Sprintf("file:%s/data/%s.db?%s",
		exPath,
		config.DBName,
		config.DBOption,
	)
}

func openDB(databaseUrl string) (*ent.Client, error) {
	db, err := sql.Open("sqlite3", databaseUrl)
	if err != nil {
		utils.AppLogger.Fatalf("Error opening connection to repository: %v", err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.SQLite, db)
	return ent.NewClient(ent.Driver(drv)), err
}

func NewEntClient() (*ent.Client, error) {
	client, err := openDB(getDBConnStr(config.GetUnifiedConfig()))
	if err != nil {
		utils.AppLogger.Fatalf("Failed opening connection to repository server: %v", err)
	}

	return client, err
}
