package routes

import (
	"net/http"

	"github.com/dkrest1/task-manager/controllers"
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
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/users/{userId}", userController.FindUser).Methods("GET")
	router.HandleFunc("/users/{userId}", userController.UpdateUser).Methods("PATCH")
	router.HandleFunc("/users/{userId}", userController.DeleteUser).Methods("DELETE")

	// Task
	router.HandleFunc("/users/{userId}/tasks", taskController.CreateTask).Methods("POST")
	router.HandleFunc("/users/{userId}/tasks", taskController.GetTasks).Methods("GET")
	router.HandleFunc("/users/{userId}/tasks/{taskId}", taskController.FindTask).Methods("GET")
	router.HandleFunc("/users/{userId}/tasks/{taskId}", taskController.UpdateTask).Methods("PATCH")
	router.HandleFunc("/users/{userId}/tasks/{taskId}", taskController.DeleteTask).Methods("DELETE")

	return router
}