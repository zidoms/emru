package main

import (
	"database/sql"
	"encoding/json"
	"time"

	log "github.com/limetext/log4go"
	_ "github.com/mattn/go-sqlite3"
)

type List struct {
	tasks []*task
	db    *sql.DB
}

func newList() *List {
	l := &List{tasks: make([]*task, 0)}
	l.initDB()
	return l
}

func loadList() *List {
	list := newList()

	var (
		id          int
		title, body string
		done        bool
		date        time.Time
	)

	tasks, err := list.db.Query("select * from tasks order by id desc")
	defer tasks.Close()
	if err != nil {
		panic(err)
	}
	for tasks.Next() {
		if err = tasks.Scan(&id, &title, &body, &done, &date); err != nil {
			log.Warn(err)
		}
		t := NewTask(title, body)
		t.id = id
		t.done = status(done)
		t.createdAt = date
		list.addTask(t)
	}

	return list
}

func (l *List) initDB() {
	db, err := sql.Open("sqlite3", "emru.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
		create table if not exists
			tasks(id integer primary key autoincrement,
			title, body varchar(255), done boolean default false,
			created_at datetime default current_timestamp)
	`)
	if err != nil {
		panic(err)
	}
	l.db = db
}

func (l *List) addTask(t *task) {
	log.Finest("Adding task %s", t)
	l.tasks = append(l.tasks, t)
}

func (l *List) AddTask(t *task) {
	_, err := l.db.Exec(`
		insert into tasks(title, body, done, created_at) values(?, ?, ?, ?)
		`, t.Title, t.Body, bool(t.done), t.createdAt)
	if err != nil {
		log.Error("Couldn't insert task: %s", err)
	} else {
		l.addTask(t)
	}
}

func (l *List) removeTaskByIndex(i int) {
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
	r := make([]*task, len(l.tasks))
	copy(r, l.tasks)
	return r
}

func (l *List) clear() {
	l.tasks = nil
}

func (l *List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Tasks []*task `json:"tasks"`
	}{
		l.tasks,
	})
}
