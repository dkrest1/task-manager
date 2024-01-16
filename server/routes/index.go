package routes

import (
	"net/http"

	"github.com/dkrest1/task-manager/controllers"
	"github.com/gorilla/mux"
)

func Routes(userController *controllers.UserController, taskController *controllers.TaskController) http.Handler {

	router := mux.NewRouter()

	//User
	router.HandleFunc("/users", userController.CreateUser).Methods("POST")
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")
	router.HandleFunc("/{userId}", userController.FindUser).Methods("GET")
	router.HandleFunc("/{userId}", userController.UpdateUser).Methods("PATCH")
	router.HandleFunc("/{userId}", userController.DeleteUser).Methods("DELETE")

	// Task
	router.HandleFunc("/users/{userId}/tasks", taskController.CreateTask).Methods("POST")
	router.HandleFunc("/users/{userId}/tasks", taskController.GetTasks).Methods("GET")
	router.HandleFunc("/users/{userId}/tasks/{taskId}", taskController.FindTask).Methods("GET")
	router.HandleFunc("/users/{userId}/tasks/{taskId}", taskController.UpdateTask).Methods("PATCH")
	router.HandleFunc("/users/{userId}/tasks/{taskId}", taskController.DeleteTask).Methods("DELETE")

	return router
}