package handler

import (
	"encoding/json"
	"gotaskify/models"
	"io"
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
	response, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	DB.Create(&task)
	response, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
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
	response, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "Error fetching task", http.StatusInternalServerError)
		return
	}
	w.Write(response)
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
	if task.Done {
		http.Error(w, "Task already completed", http.StatusBadRequest)
		return
	} else {
		task.Done = true
	}
	DB.Save(&task)
	jsonData, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
