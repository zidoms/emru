package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httputil"
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
		if req, err := httputil.DumpRequest(r, true); err != nil {
			log.Error("error on dumping request")
		} else {
			log.Info("%s", req)
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(h.data); err != nil {
		log.Error(err)
	}
}

func (h *ListHandler) parseReq() error {
	log.Info("parsing %s %s", h.req.Method, h.req.URL.Path)
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
	log.Info("list %s", name)
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
		log.Info("get lists")
		h.data, err = json.Marshal(h.ls)
		return
	case "POST":
		return h.newList()
	default:
		return undefinedMethodErr
	}
}

func (h *ListHandler) newList() error {
	log.Info("new list")
	var newList struct {
		Name  string       `json:"name"`
		Tasks []*emru.Task `json:"tasks"`
	}
	if err := json.NewDecoder(h.req.Body).Decode(&newList); err != nil {
		return err
	}

	if newList.Name == "" {
		return errors.New("can not use empty name")
	}
	if _, exist := h.ls[newList.Name]; exist {
		return errors.New("this name currently exists")
	}
	lst := emru.NewList()
	for _, task := range newList.Tasks {
		lst.Add(task)
	}
	h.ls[newList.Name] = lst
	return nil
}

func (h *ListHandler) listReq(name string) (err error) {
	switch h.req.Method {
	case "GET":
		log.Info("get list %s", name)
		h.data, err = json.Marshal(h.ls[name])
		return
	case "DELETE":
		return h.deleteList(name)
	default:
		return undefinedMethodErr
	}
}

func (h *ListHandler) deleteList(name string) error {
	log.Info("delete list %s", name)
	if _, exist := h.ls[name]; !exist {
		return errors.New("list doesn't exist")
	}
	delete(h.ls, name)
	return nil
}

func (h *ListHandler) tasksReq(l *emru.List) (err error) {
	switch h.req.Method {
	case "GET":
		log.Info("get tasks")
		h.data, err = json.Marshal(l.Tasks())
		return
	case "POST":
		return h.newTask(l)
	default:
		return undefinedMethodErr
	}
}

func (h *ListHandler) newTask(l *emru.List) (err error) {
	log.Info("new task")
	task := emru.NewTask("", "")
	if err = json.NewDecoder(h.req.Body).Decode(task); err != nil {
		return
	}
	h.data, err = json.Marshal(l.Add(task))
	return
}

func (h *ListHandler) taskReq(l *emru.List, id int) error {
	switch h.req.Method {
	case "GET":
		return h.task(l, id)
	case "PUT":
		return h.updateTask(l, id)
	case "DELETE":
		log.Info("remove task %d", id)
		return l.Remove(id)
	default:
		return undefinedMethodErr
	}
}

func (h *ListHandler) task(l *emru.List, id int) error {
	log.Info("get task %d", id)
	task, err := l.Get(id)
	if err != nil {
		return err
	}
	h.data, err = json.Marshal(task)
	return err
}

func (h *ListHandler) updateTask(l *emru.List, id int) error {
	log.Info("update task %d", id)
	task := emru.Task{}
	if err := json.NewDecoder(h.req.Body).Decode(&task); err != nil {
		return err
	}
	return l.Update(id, task)
}
