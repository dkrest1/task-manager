package routes

import (
	"net/http"

	"github.com/dkrest1/task-manager/controllers"
	"github.com/dkrest1/task-manager/middlewares"
	"github.com/gorilla/mux"
)

func Routes(
	userController *controllers.UserController, 
	taskController *controllers.TaskController, 
	authController *controllers.AuthController,
) http.Handler {

	router := mux.NewRouter()

	// Login
	router.HandleFunc("/auth/login", authController.Login).Methods("POST")
	router.HandleFunc("/auth/signup", userController.CreateUser).Methods("POST")

	//User
	router.HandleFunc("/users", middlewares.Auth(http.HandlerFunc(userController.GetUsers))).Methods("GET")
	router.HandleFunc("/users/{userId}", middlewares.Auth(http.HandlerFunc(userController.FindUser))).Methods("GET")
	router.HandleFunc("/users/{userId}", middlewares.Auth(http.HandlerFunc(userController.UpdateUser))).Methods("PATCH")
	router.HandleFunc("/users/{userId}", middlewares.Auth(http.HandlerFunc(userController.DeleteUser))).Methods("DELETE")

	// Task
	router.HandleFunc("/users/{userId}/tasks", middlewares.Auth(http.HandlerFunc(taskController.CreateTask))).Methods("POST")
	router.HandleFunc("/users/{userId}/tasks", middlewares.Auth(http.HandlerFunc(taskController.GetTasks))).Methods("GET")
	router.HandleFunc("/users/{userId}/tasks/{taskId}", middlewares.Auth(http.HandlerFunc(taskController.FindTask))).Methods("GET")
	router.HandleFunc("/users/{userId}/tasks/{taskId}", middlewares.Auth(http.HandlerFunc(taskController.UpdateTask))).Methods("PATCH")
	router.HandleFunc("/users/{userId}/tasks/{taskId}", middlewares.Auth(http.HandlerFunc(taskController.DeleteTask))).Methods("DELETE")

	return router
}