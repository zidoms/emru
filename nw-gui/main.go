package main

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/bmizerany/pat"
	log "github.com/limetext/log4go"
	"github.com/zidoms/emru/lists"
	"github.com/zidoms/emru/lists/tasks"
)

var list = lists.LoadList()

func main() {
	log.AddFilter("console", log.FINEST, log.NewConsoleLogWriter())

	go serve()

	_, err := exec.Command("nw", "--remote-debugging-port=9222", "./frontend/app", "4040").Output()
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
	if data, err := json.Marshal(list); err != nil {
		http.Error(w, "Couldn't marshal list", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func newTask(w http.ResponseWriter, req *http.Request) {
	log.Finest("Recieved request for new task")
	decoder := json.NewDecoder(req.Body)
	t := tasks.NewTask("", "")
	if err := decoder.Decode(t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	list.AddTask(t)
}

func updateTask(w http.ResponseWriter, req *http.Request) {
	i, err := strconv.Atoi(req.URL.Query().Get(":id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t := list.GetTask(i)
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
		list.RemoveTaskByIndex(i)
	}
}
