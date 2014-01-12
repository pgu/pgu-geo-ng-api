package api

import (
	"encoding/json"
	"net/http"
)

func init() {
	http.HandleFunc("/tasks/", handleTasks) // in a future, let's use github.com/gorilla/mux
}

func handleTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://pgu-geo-ng.appspot.com")
	w.Header().Set("Content-Type", "application/json")
	tasks := []string{"task 1", "task 2", "task 3", "task 4"}
	json.NewEncoder(w).Encode(tasks)
}
