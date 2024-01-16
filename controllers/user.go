package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/dkrest1/task-manager/configs"
	"github.com/dkrest1/task-manager/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)


type UserController struct{
	DB *gorm.DB
}

func NewUserController () *UserController {
	return &UserController{
		DB: configs.DB,
	}
}

func(c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid request Data", http.StatusBadRequest)
	}

	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" || newUser.Name == "" {
		http.Error(w, "Username, Email, and Password are required fields", http.StatusBadRequest)
	}

	if err :=  c.DB.Create(&newUser).Error; err != nil {
		http.Error(w, "Failed to create a user", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
	
}

func(c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	result := c.DB.Find(&users)

	if result.Error != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-Type", "applicatioon/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func(c *UserController) FindUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)

	userId := params["userId"]

	result := c.DB.Where("id = ?", userId).First(&user)

	if result.Error != nil {
		if errors .Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
		}else {
			http.Error(w, "Failed to find user", http.StatusInternalServerError)
		}
		
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func(c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	params := mux.Vars(r)

	userId := params["userId"]

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	result := c.DB.Where("id = ?", userId).Updates(&user)

	if result.Error != nil {
		http.Error(w, "Failed to fetch user", http.StatusInternalServerError)

		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	var updatedUser models.User

	if err := c.DB.First(&updatedUser, userId).Error; err != nil {
		http.Error(w, "Failed to fetch updated user", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

func(c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user  models.User

	params := mux.Vars(r)

	userId := params["userId"]

	result := c.DB.Where("id = ?", userId).Delete(&user)

	if result.Error != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted Successfully")
}


