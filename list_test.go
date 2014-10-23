package emru

import (
	"encoding/json"
	"testing"
)

func TestNewList(t *testing.T) {
	l := NewList()
	if len(l.tasks) != 0 {
		t.Errorf("Expected tasks be emty on new but is %d", len(l.tasks))
	}
}

func TestAddTask(t *testing.T) {
	tests := []struct {
		title string
		body  string
	}{
		{
			"Test",
			"Test Task",
		},
		{
			"Task",
			"Second Test",
		},
	}
	l := NewList()
	for i, test := range tests {
		task := NewTask(test.title, test.body)
		l.AddTask(task)
		if l.tasks[i].title != test.title {
			t.Errorf("Test %d: Expected task title %s, but got %s", test.title, l.tasks[i].title)
		}
		if l.tasks[i].body != test.body {
			t.Errorf("Test %d: Expected task body %s, but got %s", test.body, l.tasks[i].body)
		}
	}
}

func TestTasks(t *testing.T) {
	l := NewList()
	if len(l.Tasks()) != 0 {
		t.Errorf("Expected tasks be empty but has %d members", len(l.Tasks()))
	}
	task := NewTask("", "")
	l.AddTask(task)
	if len(l.Tasks()) != 1 {
		t.Errorf("Expected tasks len 0, but got %d", len(l.Tasks()))
	}
}

func TestMarshalJsonList(t *testing.T) {
	l := NewList()
	t1 := NewTask("Task Title", "Task Body")
	l.AddTask(t1)
	b, err := json.Marshal(l)
	if err != nil {
		t.Fatalf("Couldn't marshal list: %s", err)
	}
	if string(b) != `{"tasks":[{"title":"Task Title","body":"Task Body"}]}` {
		t.Errorf("Expected marshaled json %s, but got %s", string(b), `{"tasks":[{"title":"Task Title","body":"Task Body"}]}`)
	}
}
