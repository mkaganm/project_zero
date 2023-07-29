package main

import (
	"cronitor/internal/config"
	"cronitor/internal/scripts"
	"cronitor/internal/utils"
	"github.com/robfig/cron/v3"
)

func main() {

	// Init env configs
	config.InitEnvConfigs()

	// create cron
	c := cron.New()

	// Add reset login attempt counts cron job
	_, err := c.AddFunc("*/15 * * * *", scripts.ResetLoginAttemptCounts)
	utils.FatalErr("Error while adding cron job", err)

	_, err = c.AddFunc("*/15 * * * *", scripts.DeleteExpiredVerifications)
	utils.FatalErr("Error while adding cron job", err)

	_, err = c.AddFunc("0 0 * * *", scripts.CleanMongoLogs)

	c.Start()

	// Wait until the application is stopped
	select {}
}
