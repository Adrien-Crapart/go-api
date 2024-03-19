package main

import (
	middleware "cards/middlewares"
	"fmt"
	"go-api/config"
	"go-api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    // Load configurations
    config.InitConfig()

    // Connect to database
    config.InitDB()

    // Create a new router
    r := mux.NewRouter()

    // Middleware
    r.Use(middleware.LogRequest)

    // Handlers
    r.HandleFunc("/api/resource", handlers.GetAllResources).Methods("GET")
    r.HandleFunc("/api/resource/{id}", handlers.GetResource).Methods("GET")
    r.HandleFunc("/api/resource", handlers.CreateResource).Methods("POST")
    r.HandleFunc("/api/resource/{id}", handlers.UpdateResource).Methods("PUT")
    r.HandleFunc("/api/resource/{id}", handlers.DeleteResource).Methods("DELETE")

    // Start server
    port := config.AppConfig.Port
    fmt.Printf("Server running on port %s\n", port)
    http.ListenAndServe(":"+port, r)
}