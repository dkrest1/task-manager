package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/dkrest1/task-manager/configs"
	"github.com/dkrest1/task-manager/models"
	"github.com/dkrest1/task-manager/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// User response
type UserResponse struct {
	ID         uint       `json:"id"`
	Username   string     `json:"username"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`

}

func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

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
		return
	}


	if newUser.Username == "" || newUser.Email == "" || newUser.Password == "" || newUser.Name == "" {
		http.Error(w, "Username, Email, and Password are required fields", http.StatusBadRequest)
		return
	}

	var existingEmail models.User
	err = c.DB.Where("email = ?", newUser.Email).First(&existingEmail).Error
	if err == nil {
		http.Error(w, "Email already exist", http.StatusBadRequest)
		return
	}else if !errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "Failed to check email existence", http.StatusInternalServerError)
		return
	}

	var existingUsername models.User
	err = c.DB.Where("username = ?", newUser.Username).First(&existingUsername).Error
	if err == nil {
		http.Error(w, "Username already exist", http.StatusBadRequest)
		return
	}else if !errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, "Failed to check username existence", http.StatusInternalServerError)
		return
	}

	// Hash password
	hashPassword, err := utils.HashPassword(newUser.Password)

	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	newUser.Password = hashPassword

	if err :=  c.DB.Create(&newUser).Error; err != nil {
		http.Error(w, "Failed to create a user", http.StatusInternalServerError)
		return 
	}

	// User response
	userResponse := NewUserResponse(&newUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userResponse)
	
}

func(c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	result := c.DB.Find(&users)

	if result.Error != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return 
	}

	var userResponses []UserResponse
	for _, user := range users {
		userResponse := NewUserResponse(&user)
		userResponses = append(userResponses, userResponse)
	}

	w.Header().Set("Content-Type", "applicatioon/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponses)
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

	userResponse := NewUserResponse(&user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
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

	userResponse := NewUserResponse(&updatedUser)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userResponse)
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


