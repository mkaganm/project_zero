package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"userservice/internal/utils"
)

// EnvConfigs is a global variable that contains all environment variables
var EnvConfigs *envConfigs

// InitEnvConfigs initializes the EnvConfigs variable
func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

// envConfigs is a struct that contains all environment variables
type envConfigs struct {
	LocalServerPort  string `mapstructure:"LOCAL_SERVER_PORT"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPass     string `mapstructure:"POSTGRES_PASS"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	PostgresSSLMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	MailerUrl        string `mapstructure:"MAILER_URL"`
	LoggerSuccessUrl string `mapstructure:"LOGGER_SUCCESS_URL"`
	LoggerErrorUrl   string `mapstructure:"LOGGER_ERROR_URL"`
}

// loadEnvVariables loads all environment variables from the userservice.env file
func loadEnvVariables() *envConfigs {
	// Tell the viper the path/location of the configuration file
	viper.AddConfigPath(".")
	// Tell viper the name of the configuration file (without the extension)
	viper.SetConfigName("userservice")
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
