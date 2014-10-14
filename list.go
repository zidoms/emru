package emru

import (
	"sync"
	"time"
)

// Today's todo list
type list struct {
	tasks []task
	date  time.Time
	lock  sync.Mutex
}

func newList() *list {
	return &list{tasks: nil, date: time.Now()}
}

func (l *list) addTask(t task) {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.tasks = append(l.tasks, t)
}
