package handler

import "net/http"

// ...existing code...

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte("server is running successfully!"))
}
