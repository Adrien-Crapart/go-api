package handlers

import (
	"encoding/json"
	"net/http"

	"./config"
	"./models"

	"github.com/gorilla/mux"
)

func GetAllResources(w http.ResponseWriter, r *http.Request) {
    var resources []models.Resource
    config.DB.Find(&resources)
    json.NewEncoder(w).Encode(resources)
}

func GetResource(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var resource models.Resource
    config.DB.First(&resource, params["id"])
    json.NewEncoder(w).Encode(resource)
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
    var resource models.Resource
    _ = json.NewDecoder(r.Body).Decode(&resource)
    config.DB.Create(&resource)
    json.NewEncoder(w).Encode(resource)
}

func UpdateResource(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var resource models.Resource
    config.DB.First(&resource, params["id"])
    _ = json.NewDecoder(r.Body).Decode(&resource)
    config.DB.Save(&resource)
    json.NewEncoder(w).Encode(resource)
}

func DeleteResource(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var resource models.Resource
    config.DB.Delete(&resource, params["id"])
}