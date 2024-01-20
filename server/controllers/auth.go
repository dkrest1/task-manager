package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dkrest1/task-manager/configs"
	"github.com/dkrest1/task-manager/models"
	"gorm.io/gorm"
)


type AuthController struct{
	DB *gorm.DB
}

func NewAuthController () *AuthController {
	return &AuthController{
		DB: configs.DB,
	}
}

func(c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		http.Error(w, "Invalid login credentials", http.StatusBadRequest)
		return
	}

	if loginUser.Email == "" || loginUser.Password == "" {
		http.Error(w, "email and password are required", http.StatusBadRequest)
		return
	}

	var existingUser models.User

	result := c.DB.Where("email = ?", loginUser.Email).First(&existingUser)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)

		}else {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		}

		return
	}

	// Compare login password to user password


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingUser)

}

