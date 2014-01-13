package api

import (
	"encoding/json"
	"log"
	"net/http"

	"appengine"
	"github.com/gorilla/mux"
)

const TasksPrefix = "/tasks/"

func init() {
	r := mux.NewRouter()
	r.HandleFunc(TasksPrefix, wrapHandler(ListTasks)).Methods("GET")
	http.Handle(TasksPrefix, r)
}

type badRequest struct{ error }
type notFound struct{ error }

func wrapHandler(f func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// pre ops:
		// - setup response
		front_origin := ""
		if appengine.IsDevAppServer() {
			front_origin = "http://127.0.0.1:9000"
		} else {
			front_origin = "http://pgu-geo-ng.appspot.com"
		}

		w.Header().Set("Access-Control-Allow-Origin", front_origin)
		w.Header().Set("Content-Type", "application/json")

		// execute handler
		err := f(w, r)

		// post ops
		// - handle an error
		if err == nil {
			return
		}
		switch err.(type) {
		case badRequest:
			http.Error(w, err.Error(), http.StatusBadRequest)
		case notFound:
			http.Error(w, "entity not found", http.StatusNotFound)
		default:
			log.Println(err)
			http.Error(w, "oops", http.StatusInternalServerError)
		}
	}
}

func ListTasks(w http.ResponseWriter, r *http.Request) error {
	tasks := []string{"task 1", "task 2", "task 3", "task 4"}
	return json.NewEncoder(w).Encode(tasks)
}
