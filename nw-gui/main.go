package main

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/bmizerany/pat"
	log "github.com/limetext/log4go"
	"github.com/zidoms/emru/list"
	"github.com/zidoms/emru/list/task"
)

var lst = list.LoadList()

func main() {
	log.AddFilter("console", log.FINEST, log.NewConsoleLogWriter())

	go serve()

	_, err := exec.Command("nw", "--remote-debugging-port=9222", "./app", "4040").Output()
	if err != nil {
		panic(err)
	}
}

func serve() {
	r := pat.New()
	r.Get("/", http.HandlerFunc(getList))
	r.Post("/tasks", http.HandlerFunc(newTask))
	r.Put("/tasks/:id", http.HandlerFunc(updateTask))
	r.Del("/tasks/:id", http.HandlerFunc(deleteTask))
	http.Handle("/", r)
	http.ListenAndServe(":4040", nil)
}

func getList(w http.ResponseWriter, req *http.Request) {
	log.Finest("Recieved request for list")
	if data, err := json.Marshal(lst); err != nil {
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
	lst.AddTask(t)
}

func updateTask(w http.ResponseWriter, req *http.Request) {
	i, err := strconv.Atoi(req.URL.Query().Get(":id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t := lst.GetTask(i)
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func deleteTask(w http.ResponseWriter, req *http.Request) {
	if i, err := strconv.Atoi(req.URL.Query().Get(":id")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		lst.RemoveTask(i)
	}
}
