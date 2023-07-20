package utils

import (
	"go.uber.org/zap"
	"log"
)

var logger *zap.Logger

func init() {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"logfile.log"}

	var err error
	logger, err = cfg.Build()
	if err != nil {
		log.Panic("Logger could not started! | ", err)
	}
}
