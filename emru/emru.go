package main

import (
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	log "github.com/limetext/log4go"
	"github.com/zoli/emru/list"
)

type ListHandler struct {
	ls     map[string]list.List
	req    *http.Request
	writer http.ResponseWriter
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.req = r
	h.writer = w
	if err := h.parseReq(r); err != nil {
		log.Error(err)
		http.NotFound(w, r)
	}
}

func (h *ListHandler) parseReq(r *http.Request) error {
	log.Debug("Parsing %s", r.URL.Path)
	url := strings.TrimRight(r.URL.Path, "/")

	if url[:6] != "/lists" {
		return errors.New("invalid request")
	}
	if len(url) == 6 {
		switch h.req.Method {
		case "GET":
			h.lists()
		case "POST":
			h.newList()
		default:
			return errors.New("undefined method")
		}
		return nil
	}

	path := strings.Split(url[6:], "/")
	name := path[0]
	l, exist := h.ls[name]
	if !exist {
		return errors.New("list " + name + " not found")
	}
	if len(path) == 1 {
		switch h.req.Method {
		case "GET":
			h.list(name)
		case "PUT":
			h.updateList(name)
		case "DELETE":
			h.deleteList(name)
		default:
			return errors.New("undefined method")
		}
		return nil
	}

	if path[1] != "tasks" {
		return errors.New("invalid request")
	}
	if len(path) == 2 {
		switch h.req.Method {
		case "GET":
			h.tasks(l)
		case "POST":
			h.newTask(l)
		default:
			return errors.New("undefined method")
		}
		return nil
	}

	if len(path) > 3 {
		return errors.New("invalid request")
	}
	id, err := strconv.Atoi(path[2])
	if err != nil {
		return errors.New("invalid task id")
	}
	switch h.req.Method {
	case "GET":
		h.task(l, id)
	case "PUT":
		h.updateTask(l, id)
	case "DELETE":
		h.deleteTask(l, id)
	default:
		return errors.New("undefined method")
	}
	return nil
}

func (h *ListHandler) lists() {

}

func (h *ListHandler) list(name string) {

}

func (h *ListHandler) newList() {

}

func (h *ListHandler) updateList(name string) {

}

func (h *ListHandler) deleteList(name string) {

}

func (h *ListHandler) tasks(l list.List) {

}

func (h *ListHandler) task(l list.List, id int) {

}

func (h *ListHandler) newTask(l list.List) {

}

func (h *ListHandler) updateTask(l list.List, id int) {

}

func (h *ListHandler) deleteTask(l list.List, id int) {

}

func main() {
	l, err := net.Listen("unix", "/tmp/emru.sock")
	if err != nil {
		log.Critical(err)
		panic(err)
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	go func() {
		<-c
		os.Remove("/tmp/emru.sock")
		os.Exit(0)
	}()

	handler := &ListHandler{ls: make(map[string]list.List)}
	if http.Serve(l, handler); err != nil {
		log.Critical(err)
		panic(err)
	}
}
