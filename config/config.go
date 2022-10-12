package config

import (
	utils "urlshortener/utility"

	"github.com/gookit/validate"
	"github.com/spf13/viper"
)

type UnifiedConfig struct {
	AppEnv             string `mapstructure:"APP_ENV"`
	AppLogLevel        string `mapstructure:"LOG_LEVEL"`
	AppShowBanner      bool   `mapstructure:"SHOW_BANNER"`
	AppCertPath        string `mapstructure:"CERT_PATH"`
	AppCAFile          string `mapstructure:"CA_FILE"`
	AppTimeout         int    `mapstructure:"APP_TIMEOUT"`
	AppIDMeta          byte   `mapstructure:"APP_ID_META"`
	AppSystemID        string `mapstructure:"APP_SYSTEM_ID"`
	HTTPListen         string `mapstructure:"HTTP_LISTEN"`
	HTTPPort           int    `mapstructure:"HTTP_PORT"`
	DBDriver           string `mapstructure:"DB_DRIVER"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             int    `mapstructure:"DB_PORT"`
	DBName             string `mapstructure:"DB_NAME"`
	DBUser             string `mapstructure:"DB_USER"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	DBOption           string `mapstructure:"DB_OPTION"`
	MongoHost          string `mapstructure:"MONGO_HOST"`
	MongoPort          int    `mapstructure:"MONGO_PORT"`
	MongoDBName        string `mapstructure:"MONGO_DB_NAME"`
	MongoLoginDB       string `mapstructure:"MONGO_LOGIN_DB"`
	MongoUser          string `mapstructure:"MONGO_USER"`
	MongoPassword      string `mapstructure:"MONGO_PASSWORD"`
	MongoOption        string `mapstructure:"MONGO_OPTION"`
	NATSHost           string `mapstructure:"NATS_HOST"`
	NATSPort           int    `mapstructure:"NATS_PORT"`
	NATSName           string `mapstructure:"NATS_NAME"`
	NATSUser           string `mapstructure:"NATS_USER"`
	NATSPassword       string `mapstructure:"NATS_PASSWORD"`
	NATSNkeysSeed      string `mapstructure:"NATS_NKEYS_SEED"`
	NATSTls            bool   `mapstructure:"NATS_TLS"`
	KcBaseUrl          string `mapstructure:"KC_BASE_URL"`
	KcRealm            string `mapstructure:"KC_REALM"`
	KcClientId         string `mapstructure:"KC_CLIENT_ID"`
	KcClientCredential string `mapstructure:"KC_CLIENT_CREDENTIAL"`
}

var (
	//appConfig     *Config
	unifiedConfig *UnifiedConfig
	PathToSkip    = map[string]bool{
		"/":        true,
		"/version": true,
	}
)

func init() {
	utils.AppLogger.Info("Initializing Application Configuration")

	unifiedConfig = initUnifiedConfig()

	//Configure validation global options
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		opt.SkipOnEmpty = false
	})

	utils.AppLogger.Info("Application Configuration initialization complete")
}

func initUnifiedConfig() *UnifiedConfig {
	var cfgInitErr error

	viper.SetConfigFile(".env") // name of config file (without extension)
	//viper.SetConfigType("env")               // REQUIRED if the config file does not have the extension in the name
	// viper.AddConfigPath("/etc/otp-server/")  // path to look for the config file in
	// viper.AddConfigPath("$HOME/.otp-server") // call multiple times to add many search paths
	// viper.AddConfigPath(".")                 // optionally look for config in the working directory

	viper.AutomaticEnv()

	cfgInitErr = viper.ReadInConfig() // Find and read the config file
	if cfgInitErr != nil {            // Handle errors reading the config file
		utils.AppLogger.Fatalf("Fatal error while reading application configuration file: %v", cfgInitErr)
	}

	cfg := &UnifiedConfig{}
	cfgInitErr = viper.Unmarshal(cfg)
	if cfgInitErr != nil {
		utils.AppLogger.Fatalf("Unable to unmarshal configuration into struct: %v", cfgInitErr)
	}
	utils.AppLogger.Info("Application Configuration unmarshal successful")
	return cfg
}

func GetUnifiedConfig() *UnifiedConfig {
	if unifiedConfig == nil {
		utils.AppLogger.Info("Application Configuration is nil. Initializing new...")
		unifiedConfig = initUnifiedConfig()
	}

	return unifiedConfig
}
