package config

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"mailerservice/internal/utils"
)

// EnvConfigs is a global variable that contains all environment variables
var EnvConfigs *envConfigs

// init initializes the EnvConfigs variable
func init() {
	EnvConfigs = loadEnvVariables()
}

// envConfigs is a struct that contains all environment variables
type envConfigs struct {
	LocalServerPort      string `mapstructure:"LOCAL_SERVER_PORT"`
	SmtpHost             string `mapstructure:"SMTP_HOST"`
	SmtpPort             int    `mapstructure:"SMTP_PORT"`
	MailerSenderName     string `mapstructure:"MAILER_SENDER_NAME"`
	MailerSenderAddress  string `mapstructure:"MAILER_SENDER_ADDRESS"`
	MailerSenderPassword string `mapstructure:"MAILER_SENDER_PASSWORD"`
	LoggerMongoUrl       string `mapstructure:"LOGGER_MONGO_URL"`
	LoggerElasticUrl     string `mapstructure:"LOGGER_ELASTIC_URL"`
	RabbitMQHost         string `mapstructure:"RABBITMQ_HOST"`
	RabbitMQPort         string `mapstructure:"RABBITMQ_PORT"`
	RabbitMQUser         string `mapstructure:"RABBITMQ_USER"`
	RabbitMQPass         string `mapstructure:"RABBITMQ_PASS"`
}

// loadEnvVariables loads all environment variables from the userservice.env file
func loadEnvVariables() *envConfigs {
	// Tell the viper the path/location of the configuration file
	viper.AddConfigPath(".")
	// Tell viper the name of the configuration file (without the extension)
	viper.SetConfigName("mailerservice")
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
