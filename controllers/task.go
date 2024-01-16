package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dkrest1/task-manager/configs"
	"github.com/dkrest1/task-manager/models"
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
	var newTask models.Task

	err := json.NewDecoder(r.Body).Decode(&newTask)

	if err != nil {
		http.Error(w, "Invalid request Data", http.StatusBadRequest)
	}

}

func(c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {

}

func(c *TaskController) FindTask(w http.ResponseWriter, r *http.Request) {

}

func(c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {

}

func(c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {

}

