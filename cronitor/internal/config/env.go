package config

import (
	"cronitor/internal/utils"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// EnvConfigs is a global variable that contains all environment variables
var EnvConfigs *envConfigs

// init initializes the EnvConfigs variable
func init() {
	EnvConfigs = loadEnvVariables()
}

// envConfigs is a struct that contains all environment variables
type envConfigs struct {
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPass     string `mapstructure:"POSTGRES_PASS"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	PostgresSSLMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	MongoHost        string `mapstructure:"MONGO_HOST"`
	MongoPort        string `mapstructure:"MONGO_PORT"`
	MongoUser        string `mapstructure:"MONGO_USER"`
	MongoPass        string `mapstructure:"MONGO_PASS"`
	MongoDb          string `mapstructure:"MONGO_DB"`
	LoggerElasticUrl string `mapstructure:"LOGGER_ELASTIC_URL"`
	RabbitMQUser     string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPass     string `mapstructure:"RABBITMQ_PASS"`
	RabbitMQHost     string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort     string `mapstructure:"RABBITMQ_PORT"`
}

// loadEnvVariables loads all environment variables from the userservice.env file
func loadEnvVariables() *envConfigs {
	// Tell the viper the path/location of the configuration file
	viper.AddConfigPath(".")
	// Tell viper the name of the configuration file (without the extension)
	viper.SetConfigName("cronitor")
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
