package main

import "time"

// On each change that ocures to a task we will
// keep a copy of task before the change as an action
type action struct {
	before    task
	createdAt time.Time
}

func newAction(t task) *action {
	return &action{t, time.Now()}
}
