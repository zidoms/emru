package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/bmizerany/pat"
	log "github.com/limetext/log4go"
	"github.com/zoli/emru/list"
	"github.com/zoli/emru/list/task"
)

func main() {

}

func serve() {
	r := pat.New()
	r.Get("/tasks", http.HandlerFunc(tasks))
	r.Post("/tasks", http.HandlerFunc(newTask))
	r.Put("/tasks/:id", http.HandlerFunc(updateTask))
	r.Del("/tasks/:id", http.HandlerFunc(deleteTask))
	http.Handle("/", r)
	http.ListenAndServe(":4040", nil)
}

func tasks(w http.ResponseWriter, req *http.Request) {
	log.Finest("Recieved request for list")

	if data, err := json.Marshal(list.Emru().Tasks()); err != nil {
		http.Error(w, "Couldn't marshal list", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func newTask(w http.ResponseWriter, req *http.Request) {
	log.Finest("Recieved request for new task")

	decoder := json.NewDecoder(req.Body)
	t := task.NewTask("", "")
	if err := decoder.Decode(t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if data, err := json.Marshal(list.Emru().Add(t)); err != nil {
		http.Error(w, "Couldn't marshal returned task", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func updateTask(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.URL.Query().Get(":id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := task.Task{}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := list.Emru().Update(id, t); err != nil {
		log.Error("Error on updating task %d: %s", id, err)
	}
}

func deleteTask(w http.ResponseWriter, req *http.Request) {
	if id, err := strconv.Atoi(req.URL.Query().Get(":id")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		if err = list.Emru().Remove(id); err != nil {
			log.Error("Error on removing task %d: %s", id, err)
		}
	}
}
