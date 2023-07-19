package config

import (
	"cronitor/pkg/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// EnvConfigs is a global variable that contains all environment variables
var EnvConfigs *envConfigs

// InitEnvConfigs initializes the EnvConfigs variable
func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

// envConfigs is a struct that contains all environment variables
type envConfigs struct {
	PostgresHost    string `mapstructure:"POSTGRES_HOST"`
	PostgresPort    string `mapstructure:"POSTGRES_PORT"`
	PostgresUser    string `mapstructure:"POSTGRES_USER"`
	PostgresPass    string `mapstructure:"POSTGRES_PASS"`
	PostgresDb      string `mapstructure:"POSTGRES_DB"`
	PostgresSSLMode string `mapstructure:"POSTGRES_SSL_MODE"`
}

// loadEnvVariables loads all environment variables from the dev.env file
func loadEnvVariables() *envConfigs {
	// Tell the viper the path/location of the configuration file
	viper.AddConfigPath(".")
	// Tell viper the name of the configuration file (without the extension)
	viper.SetConfigName("dev")
	// Tell viper the configuration type
	viper.SetConfigType("env")

	// Read the configuration file
	err := viper.ReadInConfig()
	utils.FatalErr("Error reading config file! ", err)

	// Unmarshal the configuration file into a struct
	var config envConfigs
	err = viper.Unmarshal(&config, func(c *mapstructure.DecoderConfig) {
		c.TagName = "mapstructure"
	})
	utils.FatalErr("Unable to decode into struct! ", err)

	return &config
}
