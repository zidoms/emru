package main

import (
	"encoding/json"
	"testing"
)

func TestNewTask(t *testing.T) {
	task := NewTask("test", "body")
	if task.Title != "test" {
		t.Errorf("Expected task title %s, but got %s", "test", task.Title)
	}
	if task.Body != "body" {
		t.Errorf("Expected task body %s, but got %s", "body", task.Body)
	}
	if task.done {
		t.Errorf("Expected task status %v, but got %v", false, task.done)
	}
	if len(task.actions) != 1 {
		t.Errorf("Expected task actions len %d, but got %d", 1, len(task.actions))
	}
}

func TestString(t *testing.T) {
	task := NewTask("test", "body")
	if task.String() != "Title: test, Body: body" {
		t.Errorf("Expected task string %s, but got %s", "Title: test, Body: body", task.String())
	}
}

func TestMarshalJsonTask(t *testing.T) {
	task := NewTask("Task Title", "Task Body")
	b, err := json.Marshal(task)
	if err != nil {
		t.Fatalf("Couldn't marshal list: %s", err)
	}
	v := make(map[string]string)
	err = json.Unmarshal(b, &v)
	if err != nil {
		t.Fatalf("Couldn't unmarshal: %s", err)
	}
	if v["title"] != task.Title {
		t.Errorf("Expected title %s, but got %s", task.Title, v["title"])
	}
	if v["body"] != task.Body {
		t.Errorf("Expected body %s, but got %s", task.Body, v["body"])
	}
}

func TestToggle(t *testing.T) {
	tests := []struct {
		s      status
		expect status
	}{
		{true, false},
		{false, true},
	}
	for i, test := range tests {
		test.s.Toggle()
		if test.s != test.expect {
			t.Errorf("Test %d: Expected %s after toggle but got %s", i, test.expect, test.s)
		}
	}
}
