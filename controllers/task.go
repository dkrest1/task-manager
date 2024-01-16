package controllers

import (
	// "encoding/json"
	"fmt"
	"net/http"

	"github.com/dkrest1/task-manager/configs"
	// "github.com/dkrest1/task-manager/models"
	"gorm.io/gorm"
)


type TaskController struct{
	DB *gorm.DB
}

func NewTaskController () *TaskController {
	return &TaskController{
		DB: configs.DB,
	}
}

func(c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	  // Return a plain text response
	  w.Header().Set("Content-Type", "text/plain")
	  w.WriteHeader(http.StatusOK)
	  fmt.Fprint(w, "Task created successfully")

}

func(c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {

}

func(c *TaskController) FindTask(w http.ResponseWriter, r *http.Request) {

}

func(c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {

}

func(c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {

}

