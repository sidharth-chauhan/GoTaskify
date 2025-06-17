package main

import (
	"fmt"
	"net/http"

	"gotaskify/handler"
	"gotaskify/models"

	"github.com/gorilla/mux"
)

func main() {
	ConnectDatabase()

	// 2. AutoMigrate to create tables
	DB.AutoMigrate(&models.Task{})

	// Pass DB to handler package
	handler.InitializeDatabase(DB)

	r := mux.NewRouter()

	// Routers
	r.HandleFunc("/tasks", handler.GetAllTasks).Methods("GET")
	r.HandleFunc("/status", handler.Hello).Methods("GET") // Remove or implement this if needed
	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handler.GetTaskById).Methods("GET")
	r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")

	// Start the server on port 8080
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
