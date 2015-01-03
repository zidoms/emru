package main

import (
	"database/sql"
	"encoding/json"
	"sync"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Today's todo list
type List struct {
	tasks []*task
	date  time.Time
	db    *sql.DB
	lock  sync.Mutex
}

func newList() *List {
	return &List{tasks: make([]*task, 0), date: time.Now()}
}

// func loadList() *List {
// 	db, err := sql.Open("sqlite3", "emru.db")
// 	if err != nil {
// 		panic(err)
// 	}
// 	list := newList()
// }

func (l *List) AddTask(t *task) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.tasks = append(l.tasks, t)
}

func (l *List) removeTaskByIndex(i int) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if e := len(l.tasks) - 1; i != e {
		copy(l.tasks[i:], l.tasks[i+1:])
	} else {
		l.tasks = l.tasks[:e]
	}
}

func (l *List) getTask(i int) *task {
	return l.tasks[i]
}

func (l *List) Tasks() []*task {
	l.lock.Lock()
	defer l.lock.Unlock()
	r := make([]*task, len(l.tasks))
	copy(r, l.tasks)
	return r
}

func (l *List) clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.tasks = nil
}

func (l *List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Tasks []*task `json:"tasks"`
	}{
		l.tasks,
	})
}
