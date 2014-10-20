package main

import (
	"code.google.com/p/go.net/websocket"
	"github.com/zidoms/emru"
	"io"
	"net/http"
	"os/exec"
)

var list *emru.List

func main() {
	list = emru.NewList()

	http.Handle("/", websocket.Handler(Handle))
	go http.ListenAndServe(":4040", nil)

	_, err := exec.Command("nw", "--remote-debugging-port=9222", "./app", "4040").Output()
	if err != nil {
		panic(err)
	}
}

func Handle(ws *websocket.Conn) {
	io.Copy(ws, ws)
}
