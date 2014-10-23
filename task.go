package emru

import (
	"encoding/json"
	"fmt"
	"time"
)

type (
	task struct {
		title     string
		body      string
		done      status
		reminder  time.Time
		actions   []*action // Keeps task history
		labels    []*label
		createdAt time.Time
	}

	// Task status is it done or not
	status bool
)

func NewTask(title, body string) *task {
	t := &task{
		title:     title,
		body:      body,
		done:      false,
		actions:   make([]*action, 0),
		createdAt: time.Now(),
	}
	t.actions = append(t.actions, newAction(*t))

	return t
}

func (t *task) String() string {
	return fmt.Sprintf("Title: %s, Body: %s", t.title, t.body)
}

func (t *task) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}{
		t.title,
		t.body,
	})
}

func (s *status) Toggle() {
	*s = !(*s)
}
