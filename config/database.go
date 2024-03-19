package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
    dbConfig := AppConfig.DBConfig
    connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
        dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DBName, dbConfig.Password)

    db, err := gorm.Open("postgres", connectionString)
    if err != nil {
        panic(err)
    }

    DB = db
    DB.AutoMigrate(&Resource{})
}