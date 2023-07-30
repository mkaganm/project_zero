package data

import (
	"cronitor/internal/config"
	"cronitor/internal/utils"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// PostgresDSN is a pointer to a string that holds the data source name
var PostgresDSN *string

// InitPostgresDSN is a function that initializes the data source name
func InitPostgresDSN() {
	dsn := createDSN()
	PostgresDSN = &dsn
}

// createDSN is a function that creates a data source name
func createDSN() string {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.EnvConfigs.PostgresHost,
		config.EnvConfigs.PostgresPort,
		config.EnvConfigs.PostgresUser,
		config.EnvConfigs.PostgresPass,
		config.EnvConfigs.PostgresDb,
		config.EnvConfigs.PostgresSSLMode)

	return dsn
}

// InitPostgresDB is a function that initializes the database
func InitPostgresDB() *gorm.DB {

	dsn := *PostgresDSN

	log.Default().Println("Connecting to database...")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.FatalErr("Error while connecting to database", err)

	log.Default().Println("Connected to database.")

	return db
}

// ClosePostgresDB is a function that closes the database connection
func ClosePostgresDB(db *gorm.DB) {

	log.Default().Println("Closing database connection...")

	dbSQL, err := db.DB()
	utils.FatalErr("Error while closing the database connection", err)

	err = dbSQL.Close()
	utils.FatalErr("Error while closing the database connection", err)

	log.Default().Println("Database connection closed.")
}
