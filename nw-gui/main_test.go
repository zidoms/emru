package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/zidoms/emru/list"
	"github.com/zidoms/emru/list/task"
)

func TestGetList(t *testing.T) {
	tsk := task.NewTask("Test", "Server test")
	list.Emru().AddTask(tsk)
	defer list.Emru().Clear()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:4040", nil)
	if err != nil {
		t.Fatalf("Creating NewRequest error: %s", err)
	}
	getList(w, req)
	if w.Code != 200 {
		t.Errorf("Expected response code 200, but got %d: %v", w.Code, w.Body)
	}
	jt, _ := json.Marshal(tsk)
	exp := fmt.Sprintf(`{"tasks":[%s]}`, string(jt))
	if w.Body.String() != exp {
		t.Errorf("Expected response body %s, but got %s", exp, w.Body.String())
	}
}

func TestCreateNewTask(t *testing.T) {
	defer list.Emru().Clear()
	w := httptest.NewRecorder()
	buf := []byte(`{"title":"Test","body":"Server test"}`)
	req, err := http.NewRequest("POST", "http://localhost:4040/tasks", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatalf("Creating NewRequest error: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	newTask(w, req)
	if w.Code != 200 {
		t.Errorf("Expected response code 200, but got %d: %v", w.Code, w.Body)
	}
	ts := list.Emru().Tasks()
	if len(ts) != 1 {
		t.Fatalf("Expected 1 task in list, but got %d", len(ts))
	}
	if ts[0].String() != "Title: Test, Body: Server test" {
		t.Errorf("Expected task %s, but got %s", "Title: Test, Body: Server test", ts[0].String())
	}
}

func TestUpdateTask(t *testing.T) {
	list.Emru().AddTask(task.NewTask("Test", "Server test"))
	defer list.Emru().Clear()
	w := httptest.NewRecorder()
	buf := []byte(`{"title":"Test update","body":"Updated body"}`)
	req, err := http.NewRequest("PUT", "http://localhost:4040/tasks/0?:id=0", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatalf("Creating NewRequest error: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	updateTask(w, req)
	if w.Code != 200 {
		t.Errorf("Expected response code 200, but got %d: %v", w.Code, w.Body)
	}
	ts := list.Emru().Tasks()
	if ts[0].String() != "Title: Test update, Body: Updated body" {
		t.Errorf("Expected task %s, but got %s", "Title: Test update, Body: Updated body", ts[0].String())
	}
}

func TestDeleteTask(t *testing.T) {
	list.Emru().AddTask(task.NewTask("Test", "Server test"))
	defer list.Emru().Clear()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "http://localhost:4040/tasks/0?:id=0", nil)
	if err != nil {
		t.Fatalf("Creating NewRequest error: %s", err)
	}
	deleteTask(w, req)
	if w.Code != 200 {
		t.Errorf("Expected response code 200, but got %d: %v", w.Code, w.Body)
	}
	if ts := list.Emru().Tasks(); len(ts) != 0 {
		t.Fatalf("Expected list be empty, but has %d tasks", len(ts))
	}
}
