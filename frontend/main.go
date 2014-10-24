package main

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strconv"

	"github.com/bmizerany/pat"
	"github.com/zidoms/emru"
)

var list = emru.NewList()

func main() {
	r := pat.New()
	r.Get("/", http.HandlerFunc(getList))
	r.Post("/task", http.HandlerFunc(newTask))
	r.Put("/task/:id", http.HandlerFunc(updateTask))
	r.Del("/task/:id", http.HandlerFunc(deleteTask))
	http.Handle("/", r)
	go http.ListenAndServe(":4040", nil)

	_, err := exec.Command("nw", "--remote-debugging-port=9222", "./app", "4040").Output()
	if err != nil {
		panic(err)
	}
}

func getList(w http.ResponseWriter, req *http.Request) {
	if data, err := json.Marshal(list); err != nil {
		http.Error(w, "Couldn't marshal list", http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func newTask(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	t := emru.NewTask("", "")
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
