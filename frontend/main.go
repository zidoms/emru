package main

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"github.com/zidoms/emru"
	"net/http"
	"os/exec"
)

var list *emru.List

func main() {
	list = emru.NewList()

	http.Handle("/", websocket.Handler(wsHandler))
	go http.ListenAndServe(":4040", nil)

	_, err := exec.Command("nw", "--remote-debugging-port=9222", "./app", "4040").Output()
	if err != nil {
		panic(err)
	}
}

func wsHandler(ws *websocket.Conn) {
	var msg string
	for {
		if err := websocket.Message.Receive(ws, &msg); err != nil {
			panic(err)
			continue
		}
		res, _ := json.Marshal(list)
		websocket.Message.Send(ws, string(res))
	}
}
