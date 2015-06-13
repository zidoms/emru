package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	log "github.com/limetext/log4go"
	"github.com/zoli/emru/emru"
)

type ListHandler struct {
	ls   emru.Lists
	req  *http.Request
	data []byte
}

var (
	invalidReqErr      = errors.New("invalid request")
	undefinedMethodErr = errors.New("undefined method")
	listNotFoundErr    = errors.New("list not found")
)

func NewHandler() *ListHandler {
	return &ListHandler{ls: make(emru.Lists)}
}

func (h *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.req = r
	if err := h.parseReq(); err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(h.data); err != nil {
		log.Error(err)
	}
}

func (h *ListHandler) parseReq() error {
	log.Debug("Parsing %s %s", h.req.Method, h.req.URL.Path)
	url := strings.TrimRight(h.req.URL.Path, "/")

	if url[:6] != "/lists" {
		return invalidReqErr
	}
	if url == "/lists" {
		return h.listsReq()
	}

	path := strings.Split(url[7:], "/")
	if len(path) > 3 || url[6] != '/' {
		return invalidReqErr
	}
	name := path[0]
	l, exist := h.ls[name]
	if !exist {
		return listNotFoundErr
	}

	if len(path) == 1 {
		return h.listReq(name)
	}

	if path[1] != "tasks" {
		return invalidReqErr
	}
	if len(path) == 2 {
		return h.tasksReq(l)
	}

	id, err := strconv.Atoi(path[2])
	if err != nil {
		return undefinedMethodErr
	}
	return h.taskReq(l, id)
}

func (h *ListHandler) listsReq() (err error) {
	switch h.req.Method {
	case "GET":
		h.data, err = json.Marshal(h.ls)
		return
	case "POST":
		return h.newList()
	default:
		return undefinedMethodErr
	}
}

func (h *ListHandler) newList() error {
	decoder := json.NewDecoder(h.req.Body)
	var nlst struct {
		Name  string       `json:"name"`
		Tasks []*emru.Task `json:"tasks"`
	}
	if err := decoder.Decode(&nlst); err != nil {
		return err
	}

	if nlst.Name == "" {
		return errors.New("can not use empty name")
	}
	if _, exist := h.ls[nlst.Name]; exist {
		return errors.New("this name currently exists")
	}
	lst := emru.NewList()
	for _, tsk := range nlst.Tasks {
		lst.Add(tsk)
	}
	h.ls[nlst.Name] = lst
	return nil
}

func (h *ListHandler) listReq(name string) (err error) {
	switch h.req.Method {
	case "GET":
		h.data, err = json.Marshal(h.ls[name])
		return
	case "DELETE":
		return h.deleteList(name)
	default:
		return undefinedMethodErr
	}
}

func (h *ListHandler) deleteList(name string) error {
	if _, exist := h.ls[name]; !exist {
		return errors.New("list doesn't exist")
	}
	delete(h.ls, name)
	return nil
}

func (h *ListHandler) tasksReq(l *emru.List) (err error) {
	switch h.req.Method {
	case "GET":
		h.data, err = json.Marshal(l.Tasks())
		return
	case "POST":
		return h.newTask(l)
	default:
		return undefinedMethodErr
	}
}

func (h *ListHandler) newTask(l *emru.List) (err error) {
	decoder := json.NewDecoder(h.req.Body)
	tsk := emru.NewTask("", "")
	if err = decoder.Decode(tsk); err != nil {
		return
	}
	h.data, err = json.Marshal(l.Add(tsk))
	return
}

func (h *ListHandler) taskReq(l *emru.List, id int) error {
	switch h.req.Method {
	case "GET":
		return h.task(l, id)
	case "PUT":
		return h.updateTask(l, id)
	case "DELETE":
		return l.Remove(id)
	default:
		return undefinedMethodErr
	}
}

func (h *ListHandler) task(l *emru.List, id int) error {
	task, err := l.Get(id)
	if err != nil {
		return err
	}
	h.data, err = json.Marshal(task)
	return err
}

func (h *ListHandler) updateTask(l *emru.List, id int) error {
	tsk := emru.Task{}
	decoder := json.NewDecoder(h.req.Body)
	if err := decoder.Decode(&tsk); err != nil {
		return err
	}
	return l.Update(id, tsk)
}
