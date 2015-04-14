package list

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/zoli/emru/list/task"
)

func TestNewList(t *testing.T) {
	l := New()
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
	l := New()
	for i, test := range tests {
		tsk := task.New(test.title, test.body)
		l.add(tsk)
		if l.tasks[i].Title != test.title {
			t.Errorf("Test %d: Expected task title %s, but got %s", test.title, l.tasks[i].Title)
		}
		if l.tasks[i].Body != test.body {
			t.Errorf("Test %d: Expected task body %s, but got %s", test.body, l.tasks[i].Body)
		}
	}
}

func TestRemove(t *testing.T) {
	l := New()
	t1 := task.New("Task title", "Task Body")
	l.add(t1)
	l.remove(0)
	if len(l.Tasks()) != 0 {
		t.Errorf("Expected tasks be empty but has %d members", len(l.Tasks()))
	}
}

func TestTasks(t *testing.T) {
	l := New()
	if ts := l.Tasks(); len(ts) != 0 {
		t.Errorf("Expected tasks be empty but has %d members: %v", len(ts), ts)
	}
	task := task.New("", "")
	l.add(task)
	if ts := l.Tasks(); len(ts) != 1 {
		t.Errorf("Expected tasks len 1, but got %d: %v", len(ts), ts)
	}
}

func TestClearList(t *testing.T) {
	l := New()
	t1 := task.New("Task title", "Task Body")
	l.add(t1)
	l.clear()
	if len(l.Tasks()) != 0 {
		t.Errorf("Expected tasks be empty but has %d members", len(l.Tasks()))
	}
}

func TestMarshalJsonList(t *testing.T) {
	l := New()
	t1 := task.New("Task title", "Task Body")
	l.add(t1)
	b, err := json.Marshal(l)
	if err != nil {
		t.Fatalf("Couldn't marshal list: %s", err)
	}
	jt, _ := json.Marshal(t1)
	time, _ := l.CreatedAt.MarshalJSON()
	exp := fmt.Sprintf(`{"tasks":[%s],"created_at":%s}`, string(jt), string(time))
	if string(b) != exp {
		t.Errorf("Expected marshaled json %s, but got %s", exp, string(b))
	}
}
