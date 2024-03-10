package config

import (
	"os"
)

type Config struct {
    DBUsername string
    DBPassword string
    DBHost     string
    DBPort     string
    DBName     string
}

var AppConfig Config

func LoadConfig() {
    // Load configuration from environment variables or config files
    AppConfig = Config{
        DBUsername: os.Getenv("DB_USERNAME"),
        DBPassword: os.Getenv("DB_PASSWORD"),
        DBHost:     os.Getenv("DB_HOST"),
        DBPort:     os.Getenv("DB_PORT"),
        DBName:     os.Getenv("DB_NAME"),
    }
}

func GetDBConnectionString() string {
    // Construct and return the database connection string
    return "postgres://" + AppConfig.DBUsername + ":" + AppConfig.DBPassword + "@" + AppConfig.DBHost + ":" + AppConfig.DBPort + "/" + AppConfig.DBName + "?sslmode=disable"
}