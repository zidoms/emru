package emru

import (
	"sync"
	"time"
)

// Today's todo list
type List struct {
	tasks []task
	date  time.Time
	lock  sync.Mutex
}

func NewList() *List {
	return &List{tasks: nil, date: time.Now()}
}

func (l *List) AddTask(t task) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.tasks = append(l.tasks, t)
}
