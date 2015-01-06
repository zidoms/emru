package list

import (
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/zidoms/emru/list/task"
)

func TestnewList(t *testing.T) {
	l := newList()
	if len(l.tasks) != 0 {
		t.Errorf("Expected tasks be emty on new but is %d", len(l.tasks))
	}
}

func TestAdd(t *testing.T) {
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
	l := newList()
	for i, test := range tests {
		task := NewTask(test.title, test.body)
		l.add(task)
		if l.tasks[i].Title != test.title {
			t.Errorf("Test %d: Expected task title %s, but got %s", test.title, l.tasks[i].Title)
		}
		if l.tasks[i].Body != test.body {
			t.Errorf("Test %d: Expected task body %s, but got %s", test.body, l.tasks[i].Body)
		}
	}
}

func TestRemove(t *testing.T) {
	l := newList()
	t1 := NewTask("Task title", "Task Body")
	l.add(t1)
	l.remove(0)
	if len(l.Tasks()) != 0 {
		t.Errorf("Expected tasks be empty but has %d members", len(l.Tasks()))
	}
}

func TestTasks(t *testing.T) {
	l := newList()
	if len(l.Tasks()) != 0 {
		t.Errorf("Expected tasks be empty but has %d members", len(l.Tasks()))
	}
	task := NewTask("", "")
	l.add(task)
	if len(l.Tasks()) != 1 {
		t.Errorf("Expected tasks len 0, but got %d", len(l.Tasks()))
	}
}

func TestClearList(t *testing.T) {
	l := newList()
	t1 := NewTask("Task title", "Task Body")
	l.add(t1)
	l.clear()
	if len(l.Tasks()) != 0 {
		t.Errorf("Expected tasks be empty but has %d members", len(l.Tasks()))
	}
}

func TestMarshalJsonList(t *testing.T) {
	l := newList()
	t1 := NewTask("Task title", "Task Body")
	l.add(t1)
	b, err := json.Marshal(l)
	if err != nil {
		t.Fatalf("Couldn't marshal list: %s", err)
	}
	jt, _ := json.Marshal(t1)
	exp := fmt.Sprintf(`{"tasks":[%s]}`, string(jt))
	if string(b) != exp {
		t.Errorf("Expected marshaled json %s, but got %s", exp, string(b))
	}
}
