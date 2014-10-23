package main

import (
	"bytes"
	"github.com/zidoms/emru"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetList(t *testing.T) {
	list.AddTask(emru.NewTask("Test", "Server test"))
	defer list.Clear()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "http://localhost:4040", nil)
	if err != nil {
		t.Fatalf("Creating NewRequest error: %s", err)
	}
	getList(w, req)
	if w.Code != 200 {
		t.Errorf("Expected response code 200, but got %d", w.Code)
	}
	if w.Body.String() != `{"Tasks":[{"Title":"Test","Body":"Server test"}]}` {
		t.Errorf("Expected response body %s, but got %s", `{"Tasks":[{"Title":"Test","Body":"Server test"}]}`, w.Body.String())
	}
}

func TestNewTask(t *testing.T) {
	defer list.Clear()
	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"Title":"Test","Body":"Server test"}`)
	req, err := http.NewRequest("POST", "http://localhost:4040/task", body)
	if err != nil {
		t.Fatalf("Creating NewRequest error: %s", err)
	}
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
	list.AddTask(emru.NewTask("Test", "Server test"))
	defer list.Clear()
	w := httptest.NewRecorder()
	body := bytes.NewBufferString(`{"Title":"Test update","Body":"updated body"}`)
	req, err := http.NewRequest("POST", "http://localhost:4040/task/1", body)
	if err != nil {
		t.Fatalf("Creating NewRequest error: %s", err)
	}
	updateTask(w, req)
	if w.Code != 200 {
		t.Errorf("Expected response code 200, but got %d", w.Code)
	}
	ts := list.Tasks()
	if ts[0].String() != "Title: Test update, Body: updated body" {
		t.Errorf("Expected task %s, but got %s", "Title: Test update, Body: updated body", ts[0].String())
	}
}
