package main

import (
	"fmt"
	"gotaskify/handler"
	"gotaskify/models"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "data/todos.db" // Use relative path for local dev
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	db.AutoMigrate(&models.Task{})

	// Pass the db connection to the handler package
	handler.InitializeDatabase(db)
	r := mux.NewRouter()

	// Routers
	r.HandleFunc("/tasks", handler.GetAllTasks).Methods("GET")
	r.HandleFunc("/status", handler.Hello).Methods("GET")
	r.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", handler.GetTaskById).Methods("GET")
	r.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")

	// Start the server on port 8080
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
