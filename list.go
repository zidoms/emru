package emru

import "time"

// Today's todo list
type list struct {
	tasks []*task
	date  time.Time
}

func newList() *list {
	return &list{nil, time.Now()}
}
