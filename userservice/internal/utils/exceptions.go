package utils

import (
	"go.uber.org/zap"
	"log"
)

// FatalErr is a function that logs a fatal error
func FatalErr(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}

// LogErr is a function that logs an error
func LogErr(msg string, err error) {
	if err != nil {
		logger.Error(msg, zap.Error(err))
	}
}

// LogInfo is a function that logs an info
func LogInfo(msg string) {
	logger.Info(msg)
}
