package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/dkrest1/task-manager/configs"
	"github.com/dkrest1/task-manager/models"
	"github.com/dkrest1/task-manager/utils"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)


type LoginResponse struct {
	User   UserResponse `json:"user"`
	Token  string      `json:"token"`
}
type AuthController struct{
	DB *gorm.DB
}

func NewAuthController () *AuthController {
	//load env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}


	return &AuthController{
		DB: configs.DB,
	}
}

func(c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var loginUser models.User

	if err := json.NewDecoder(r.Body).Decode(&loginUser); err != nil {
		utils.HandleGenericResponse(w, "invalid login credentials", http.StatusBadRequest)
		return
	}

	if loginUser.Email == "" || loginUser.Password == "" {
		utils.HandleGenericResponse(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	var existingUser models.User

	result := c.DB.Where("email = ?", loginUser.Email).First(&existingUser)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			utils.HandleGenericResponse(w, "User not found", http.StatusBadRequest)

		}else {
			utils.HandleGenericResponse(w, "Failed to fetch user", http.StatusInternalServerError)
		}

		return
	}

	// Compare login password to user password
	err := utils.ComparePasswords(existingUser.Password, loginUser.Password)

	if err != nil {
		utils.HandleGenericResponse(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	secretKey, exist := os.LookupEnv("JWT_SECRET_KEY")

	if !exist {
		log.Fatal("JWT_SECRET_KEY not set in env")
	}

	token, err := utils.GenerateToken(existingUser.ID, existingUser.Email, secretKey)

	if err != nil {
		log.Fatal(err)
	}

	userResponse := NewUserResponse(&existingUser)

	loginResponse := LoginResponse{
		Token: token,
		User: userResponse,
	}

	response := utils.NewGenericResponse(http.StatusOK, "Success", loginResponse)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

