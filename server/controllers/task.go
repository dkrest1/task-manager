package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/dkrest1/task-manager/configs"
	"github.com/dkrest1/task-manager/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type TaskResponse struct {
	ID            uint        `json:"id"`
	UserId        uint        `json:"userId"`
	User          UserResponse `json:"user"`
	Title         string	  `json:"title"`
	Description   string      `json:"description"`
	DueDate       time.Time   `json:"dueDate"`
	Completed     bool        `json:"completed"`
	CreatedAt     time.Time   `json:"createdAt"`
	UpdatedAt     time.Time   `json:"updatedAt"`
}

func newTaskResponse(task *models.Task) TaskResponse {
	return TaskResponse {
		ID:           task.ID,
		UserId:       task.UserId,
		User:         NewUserResponse(&task.User),
		Title:        task.Title,
		Description:  task.Description,
		DueDate:      task.DueDate,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		UpdatedAt:    task.UpdatedAt, 
	}
} 

type TaskController struct{
	DB *gorm.DB
}

func NewTaskController () *TaskController {
	return &TaskController{
		DB: configs.DB,
	}
}

func(c *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userIdStr := params["userId"]

	userId, err := strconv.ParseUint(userIdStr, 10, 64) 

	if err != nil {
		http.Error(w, "Invalid userId format", http.StatusBadRequest)
		return 
	}

	var user models.User

	result := c.DB.Where("id = ?", userId).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
		}else {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		}

		return
	}

	var newTask  models.Task

	if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return 
	}

	if newTask.Description == "" || newTask.Title == "" {
		http.Error(w, "Task must have title and description", http.StatusBadRequest)
		return
	}

	newTask.UserId = uint(userId)

	if err := c.DB.Create(&newTask).Error; err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	c.DB.Preload("User").First(&newTask, newTask.ID)

	taskResponse := newTaskResponse(&newTask)


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(taskResponse)

}

func(c *TaskController) GetTasks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userIdStr := params["userId"]

	userId, err := strconv.ParseUint(userIdStr, 10, 64) 

	if err != nil {
		http.Error(w, "Invalid userId format", http.StatusBadRequest)
		return 
	}

	var user models.User

	result := c.DB.Where("id = ?", userId).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
		}else {
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
		}

		return
	}


	var tasks []models.Task

	if err := c.DB.Where("user_id = ?", userId).Preload("User").Find(&tasks).Error; err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return 
	}

	var taskResponses []TaskResponse

	for _, task := range tasks {
		taskResponse := newTaskResponse(&task)
		taskResponses = append(taskResponses, taskResponse)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskResponses)
}

func(c *TaskController) FindTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userIdStr := params["userId"]
	taskIdStr := params["taskId"]

	userId, err := strconv.ParseUint(userIdStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid userId format", http.StatusBadRequest)
		return
	}

	taskId, err := strconv.ParseUint(taskIdStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid taskId format", http.StatusBadRequest)
		return 
	}

	var task models.Task

	result := c.DB.Where("id = ? AND user_id = ?", taskId, userId).Preload("User").First(&task)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Task not found", http.StatusNotFound)
		}else {
			http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
		}
		return
	}

	taskResponse := newTaskResponse(&task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskResponse)
}

func(c *TaskController) UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userIdStr := params["userId"]
	taskIdStr := params["taskId"]

	userId, err := strconv.ParseUint(userIdStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid userId format", http.StatusBadRequest)
		return
	}

	taskId, err := strconv.ParseUint(taskIdStr, 10, 64)

	if err != nil {
		http.Error(w, "Invalid taskId format", http.StatusBadRequest)
		return 
	}

	var task models.Task

	result := c.DB.Where("id = ? AND user_id = ?", taskId, userId).Preload("User").First(&task)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Task not found", http.StatusNotFound)
		}else {	
			http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
		}

		return
	}

	var updatedTask models.Task

	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if updatedTask.Title != "" {
		task.Title = updatedTask.Title
	}

	if updatedTask.Description != "" {
		task.Description = updatedTask.Description
	}

	if !updatedTask.DueDate.IsZero() {
		task.DueDate = updatedTask.DueDate
	}

	if updatedTask.Completed != task.Completed {
		task.Completed = updatedTask.Completed
	}

	if err := c.DB.Save(&task).Error; err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return 
	}

	taskResponse := newTaskResponse(&task)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskResponse)
}

func(c *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userIdStr := params["userId"]
	taskIdStr := params["taskId"]

	userId, err := strconv.ParseUint(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid userId format", http.StatusBadRequest)
		return
	}

	taskId, err := strconv.ParseUint(taskIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid taskId format", http.StatusBadRequest)
		return
	}

	var task models.Task

	result := c.DB.Where("id = ? AND user_id = ?", taskId, userId).First(&task)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
		}
		return
	}

	if err := c.DB.Delete(&task).Error; err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Task deleted successfully")

}

