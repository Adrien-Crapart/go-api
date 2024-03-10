package db

import (
	"database/sql"
	"go-api/config"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
    // Initialize database connection
    connStr := config.GetDBConnectionString()
    database, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }

    // Test the connection
    err = database.Ping()
    if err != nil {
        panic(err)
    }

    db = database
}

func InsertUser(user User) error {
    // Insert user into the database
    // Example:
    _, err := db.Exec("INSERT INTO users (username, email, password) VALUES ($1, $2, $3)", user.Username, user.Email, user.Password)
    if err != nil {
        return err
    }
    return nil
}