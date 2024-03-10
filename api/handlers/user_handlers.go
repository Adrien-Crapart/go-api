package handlers

import (
	"encoding/json"
	"go-api/db"
	"net/http"
)

type User struct {
    // Define your user struct fields here
}

// RegisterHandler handles user registration
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Validate and insert user into the database
    err = db.InsertUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with success message
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("User registered successfully"))
}