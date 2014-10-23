package main

import (
	// "encoding/json"
	"net/http"
	"os/exec"

	"github.com/bmizerany/pat"
	"github.com/zidoms/emru"
)

var list = emru.NewList()

func main() {
	r := pat.New()
	r.Get("/", http.HandlerFunc(getList))
	r.Post("task", http.HandlerFunc(newTask))
	r.Post("task/:id", http.HandlerFunc(updateTask))
	http.Handle("/", r)
	go http.ListenAndServe(":4040", nil)

	_, err := exec.Command("nw", "--remote-debugging-port=9222", "./app", "4040").Output()
	if err != nil {
		panic(err)
	}
}

func getList(w http.ResponseWriter, req *http.Request) {

}

func newTask(w http.ResponseWriter, req *http.Request) {

}

func updateTask(w http.ResponseWriter, req *http.Request) {

}
