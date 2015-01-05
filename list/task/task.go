package task

import (
	"fmt"
	"time"
)

type (
	Task struct {
		Id        int       `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		Done      Status    `json:"done"`
		CreatedAt time.Time `json:"created_at"`
	}

	Tasks []*Task

	// Task status is it done or not
	Status bool
)

func NewTask(title, body string) *Task {
	t := &Task{
		Title:     title,
		Body:      body,
		Done:      false,
		CreatedAt: time.Now(),
	}

	return t
}

func (t *Task) String() string {
	return fmt.Sprintf("Title: %s, Body: %s", t.Title, t.Body)
}

func (s *Status) toggle() {
	*s = !(*s)
}

func (t Tasks) Len() int {
	return len(t)
}

func (t Tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Tasks) Less(i, j int) bool {
	return t[i].Id < t[j].Id
}
