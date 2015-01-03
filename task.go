package main

import (
	"fmt"
	"time"
)

type (
	task struct {
		id        int
		Title     string `json:"title"`
		Body      string `json:"body"`
		done      status
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
		createdAt: time.Now(),
	}

	return t
}

func (t *task) String() string {
	return fmt.Sprintf("Title: %s, Body: %s", t.Title, t.Body)
}

func (s *status) toggle() {
	*s = !(*s)
}
