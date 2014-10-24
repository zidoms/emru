package emru

import (
	"fmt"
	"time"
)

type (
	task struct {
		Title     string `json:"Title"`
		Body      string `json:"Body"`
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
		Title:     title,
		Body:      body,
		done:      false,
		actions:   make([]*action, 0),
		createdAt: time.Now(),
	}
	t.actions = append(t.actions, newAction(*t))

	return t
}

func (t *task) String() string {
	return fmt.Sprintf("Title: %s, Body: %s", t.Title, t.Body)
}

func (s *status) Toggle() {
	*s = !(*s)
}
