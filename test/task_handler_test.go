package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotaskify/handler"
	"gotaskify/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var router *mux.Router

func TestMain(m *testing.M) {
	// Setup in-memory DB with silent logger
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	db.AutoMigrate(&models.Task{})
	handler.InitializeDatabase(db)

	// Setup router
	router = mux.NewRouter()
	router.HandleFunc("/tasks", handler.GetAllTasks).Methods("GET")
	router.HandleFunc("/tasks", handler.CreateTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", handler.GetTaskById).Methods("GET")
	router.HandleFunc("/tasks/{id}", handler.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", handler.UpdateTask).Methods("PUT")

	code := m.Run()
	os.Exit(code)
}

func createTaskHelper(t *testing.T, title string) models.Task {
	task := models.Task{Title: title}
	body, _ := json.Marshal(task)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated {
		t.Fatalf("Failed to create task: %v", rr.Body.String())
	}
	var created models.Task
	json.Unmarshal(rr.Body.Bytes(), &created)
	return created
}

func TestCreateTask(t *testing.T) {
	task := models.Task{Title: "Test Task"}
	body, _ := json.Marshal(task)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status 201, got %d", rr.Code)
	}
	var created models.Task
	json.Unmarshal(rr.Body.Bytes(), &created)
	if created.Title != task.Title {
		t.Errorf("Expected title %q, got %q", task.Title, created.Title)
	}
}

func TestCreateTask_InvalidJSON(t *testing.T) {
	body := []byte(`{invalid json}`)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for invalid JSON, got %d", rr.Code)
	}
}

func TestCreateTask_EmptyTitle(t *testing.T) {
	task := models.Task{Title: ""}
	body, _ := json.Marshal(task)
	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	// The handler currently allows empty titles, so expect 201.
	// If you want to enforce non-empty titles, change the handler and update this test.
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status 201 for empty title, got %d", rr.Code)
	}
}

func TestGetAllTasks(t *testing.T) {
	createTaskHelper(t, "Task 1")
	createTaskHelper(t, "Task 2")
	req := httptest.NewRequest("GET", "/tasks", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
	var tasks []models.Task
	json.Unmarshal(rr.Body.Bytes(), &tasks)
	if len(tasks) < 2 {
		t.Errorf("Expected at least 2 tasks, got %d", len(tasks))
	}
}

func TestGetTaskById(t *testing.T) {
	task := createTaskHelper(t, "Find Me")
	url := "/tasks/" + toStr(task.ID)
	req := httptest.NewRequest("GET", url, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
	var fetched models.Task
	json.Unmarshal(rr.Body.Bytes(), &fetched)
	if fetched.Title != "Find Me" {
		t.Errorf("Expected title 'Find Me', got %q", fetched.Title)
	}
}

func TestGetTaskById_NotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/tasks/99999", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for non-existent task, got %d", rr.Code)
	}
}

func TestUpdateTask(t *testing.T) {
	task := createTaskHelper(t, "To Complete")
	url := "/tasks/" + toStr(task.ID)
	req := httptest.NewRequest("PUT", url, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
	var updated models.Task
	json.Unmarshal(rr.Body.Bytes(), &updated)
	if !updated.Done {
		t.Errorf("Expected Done=true, got false")
	}
	// Try updating again, should fail
	rr2 := httptest.NewRecorder()
	router.ServeHTTP(rr2, req)
	if rr2.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for already completed, got %d", rr2.Code)
	}
}

func TestUpdateTask_NotFound(t *testing.T) {
	req := httptest.NewRequest("PUT", "/tasks/99999", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for non-existent task, got %d", rr.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	task := createTaskHelper(t, "To Delete")
	url := "/tasks/" + toStr(task.ID)
	req := httptest.NewRequest("DELETE", url, nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}
	// Try fetching deleted task
	req2 := httptest.NewRequest("GET", url, nil)
	rr2 := httptest.NewRecorder()
	router.ServeHTTP(rr2, req2)
	if rr2.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for deleted task, got %d", rr2.Code)
	}
}

func TestDeleteTask_NotFound(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/tasks/99999", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for non-existent task, got %d", rr.Code)
	}
}

// Helper to convert uint to string
func toStr(id uint) string {
	return fmt.Sprintf("%d", id)
}
