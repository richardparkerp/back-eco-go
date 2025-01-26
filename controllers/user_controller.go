package controllers

import (
	"encoding/json"
	"net/http"

	"back-eco-go/config"
	"back-eco-go/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Not valid data", http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(w, "Error addin user to db", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		http.Error(w, "Eror getting users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
