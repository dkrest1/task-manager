package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dkrest1/task-manager/models"
	// "github.com/dkrest1/task-manager/models"
)


type UserController struct{}

func NewUserController () *UserController {
	return &UserController{}
}

func(c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, "Invalid request Data", http.StatusBadRequest)
	}

	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" {
		http.Error(w, "Username, Email, and Password are required fields", http.StatusBadRequest)
	}

	createdUser, err := c.CreateUser(newUser)

	if err != nil {
		http.Error(w, "Failed to create a user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
	
}

func(c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {

}

func(c *UserController) FindUser(w http.ResponseWriter, r *http.Request) {

}

func(c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func(c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

