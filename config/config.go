package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppConfig *Config

type Config struct {
    Port     string
    DBConfig DBConfig
}

type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func InitConfig() {
    LoadEnv()

    AppConfig = &Config{
        Port: os.Getenv("PORT"),
        DBConfig: DBConfig{
            Host:     os.Getenv("DB_HOST"),
            Port:     os.Getenv("DB_PORT"),
            User:     os.Getenv("DB_USER"),
            Password: os.Getenv("DB_PASSWORD"),
            DBName:   os.Getenv("DB_NAME"),
        },
    }
}