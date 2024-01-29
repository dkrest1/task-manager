package middlewares

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dkrest1/task-manager/configs"
	"github.com/dkrest1/task-manager/models"
	"github.com/dkrest1/task-manager/utils"
)


func Auth(next http.HandlerFunc) http.HandlerFunc {
	secretKey, exist := os.LookupEnv("JWT_SECRET_KEY")

	if !exist {
		log.Fatal("JWT_SECRET_KEY not set in env")
	}


	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			utils.HandleGenericResponse(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		} 

		splitToken := strings.Split(tokenString, "Bearer ")

		if len(splitToken) != 2 {
			utils.HandleGenericResponse(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		token := splitToken[1]

		claims, err := utils.ValidateToken(token, secretKey)

		if err != nil {
			utils.HandleGenericResponse(w, "Invalid token, please login again", http.StatusUnauthorized)
			return
		}

		Float64UserId := claims["userId"].(float64)
		userId := uint(Float64UserId)
		email := claims["email"].(string)

		var user models.User

		result := configs.DB.Where("email = ?", email).First(&user)

		if result.Error != nil {
			log.Printf("Error fetching user from the database: %v", result.Error)
			utils.HandleGenericResponse(w, "Unauthorized - User not found", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userId)
		ctx = context.WithValue(ctx, "email", email)

		next.ServeHTTP(w, r.WithContext(ctx))


	})
}