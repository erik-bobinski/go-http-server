package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/erik-bobinski/go-http-server/internal/taskstore"
)

// our server wraps the taskstore type; thread-safe
type TaskServer struct {
	store *taskstore.TaskStore
}

// constructor for type TaskServer
func NewTaskServer() *TaskServer {
	taskStore := taskstore.NewTaskStore()
	return &TaskServer{store: taskStore}
}

// helper fn: marshall any value to json, server response is the json or 500 error
func renderJSON(res http.ResponseWriter, value any) {
	json, err := json.Marshal(value)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(json)
}

func (ts *TaskServer) createTaskHandler(req *http.Request, res http.ResponseWriter) {
	log.Printf("handling task create at %s\n", req.URL.Path)
}
