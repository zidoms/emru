package emru

import (
	"encoding/json"
	"sync"
	"time"
)

// Today's todo list
type List struct {
	tasks []*task
	date  time.Time
	lock  sync.Mutex
}

func NewList() *List {
	return &List{tasks: make([]*task, 0), date: time.Now()}
}

func (l *List) AddTask(t *task) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.tasks = append(l.tasks, t)
}

func (l *List) Tasks() []*task {
	l.lock.Lock()
	defer l.lock.Unlock()
	r := make([]*task, len(l.tasks))
	copy(r, l.tasks)
	return r
}

func (l *List) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.tasks = nil
}

func (l *List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Tasks []*task `json:"Tasks"`
	}{
		l.tasks,
	})
}
