package handler

import (
	"encoding/json"
	"gotaskify/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase(db *gorm.DB) {
	DB = db
}

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	DB.Find(&tasks)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	DB.Create(&task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var task models.Task
	err := DB.First(&task, id).Error
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var task models.Task
	err := DB.First(&task, id).Error
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	DB.Delete(&task)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Task deleted successfully"))
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var task models.Task
	err := DB.First(&task, id).Error
	if err != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	DB.Save(&task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}
