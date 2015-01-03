package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetList(t *testing.T) {
	list.AddTask(NewTask("Test", "Server test"))
	defer list.clear()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:4040", nil)
	if err != nil {
		t.Fatalf("Creating NewRequest error: %s", err)
	}
	getList(w, req)
	if w.Code != 200 {
		t.Errorf("Expected response code 200, but got %d", w.Code)
	}
	if w.Body.String() != `{"tasks":[{"title":"Test","body":"Server test"}]}` {
		t.Errorf("Expected response body %s, but got %s", `{"tasks":[{"title":"Test","body":"Server test"}]}`, w.Body.String())
	}
}

func TestCreateNewTask(t *testing.T) {
	defer list.clear()
	w := httptest.NewRecorder()
	buf := []byte(`{"title":"Test","body":"Server test"}`)
	req, err := http.NewRequest("POST", "http://localhost:4040/tasks", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatalf("Creating NewRequest error: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	newTask(w, req)
	if w.Code != 200 {
		t.Errorf("Expected response code 200, but got %d", w.Code)
	}
	ts := list.Tasks()
	if len(ts) != 1 {
		t.Fatalf("Expected 1 task in list, but got %d", len(ts))
	}
	if ts[0].String() != "Title: Test, Body: Server test" {
		t.Errorf("Expected task %s, but got %s", "Title: Test, Body: Server test", ts[0].String())
	}
}

func TestUpdateTask(t *testing.T) {
	list.AddTask(NewTask("Test", "Server test"))
	defer list.clear()
	w := httptest.NewRecorder()
	buf := []byte(`{"title":"Test update","body":"Updated body"}`)
	req, err := http.NewRequest("PUT", "http://localhost:4040/tasks/0?:id=0", bytes.NewBuffer(buf))
	if err != nil {
		t.Fatalf("Creating NewRequest error: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")
	updateTask(w, req)
	if w.Code != 200 {
		t.Errorf("Expected response code 200, but got %d", w.Code)
	}
	ts := list.Tasks()
	if ts[0].String() != "Title: Test update, Body: Updated body" {
		t.Errorf("Expected task %s, but got %s", "Title: Test update, Body: Updated body", ts[0].String())
	}
}

func TestDeleteTask(t *testing.T) {
	list.AddTask(NewTask("Test", "Server test"))
	defer list.clear()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", "http://localhost:4040/tasks/0?:id=0", nil)
	if err != nil {
		t.Fatalf("Creating NewRequest error: %s", err)
	}
	deleteTask(w, req)
	if w.Code != 200 {
		t.Errorf("Expected response code 200, but got %d", w.Code)
	}
	if ts := list.Tasks(); len(ts) != 0 {
		t.Fatalf("Expected list be empty, but has %d tasks", len(ts))
	}
}
