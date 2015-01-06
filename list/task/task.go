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

func (s *Status) Val() bool {
	return *s == true
}

func (ts Tasks) Len() int {
	return len(ts)
}

func (ts Tasks) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}

func (ts Tasks) Less(i, j int) bool {
	return ts[i].Id < ts[j].Id
}

func (ts Tasks) Exists(id int) bool {
	for _, t := range ts {
		if t.Id == id {
			return true
		}
	}
	return false
}

func (ts Tasks) Index(id int) int {
	for i, t := range ts {
		if t.Id == id {
			return i
		}
	}
	return -1
}
