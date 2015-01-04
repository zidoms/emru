package lists

import (
	"database/sql"
	"encoding/json"
	"time"

	log "github.com/limetext/log4go"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/zidoms/emru/lists/tasks"
)

type List struct {
	tasks []*Task
	db    *sql.DB
}

func NewList() *List {
	l := &List{tasks: make([]*Task, 0)}
	return l
}

func LoadList() *List {
	list := NewList()
	list.initDB()

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
		t.Id = id
		t.Done = Status(done)
		t.CreatedAt = date
		list.addTask(t)
	}

	return list
}

func (l *List) initDB() {
	db, err := sql.Open("sqlite3", "emru.db")
	if err != nil {
		panic(err)
	}
	q := `create table if not exists
		tasks(id integer primary key autoincrement,
		title, body varchar(255), done boolean default false,
		created_at datetime default current_timestamp)`
	if _, err = db.Exec(q); err != nil {
		panic(err)
	}
	l.db = db
}

func (l *List) addTask(t *Task) {
	log.Finest("Adding task %s", t)
	l.tasks = append(l.tasks, t)
}

func (l *List) AddTask(t *Task) {
	q := "insert into tasks(title, body, done, created_at) values(?, ?, ?, ?)"
	if _, err := l.db.Exec(q, t.Title, t.Body, bool(t.Done), t.CreatedAt); err != nil {
		log.Error("Couldn't insert task: %s", err)
		return
	}
	l.addTask(t)
}

func (l *List) RemoveTaskByIndex(i int) {
	if e := len(l.tasks) - 1; i != e {
		copy(l.tasks[i:], l.tasks[i+1:])
	} else {
		l.tasks = l.tasks[:e]
	}
}

func (l *List) GetTask(i int) *Task {
	return l.tasks[i]
}

func (l *List) Tasks() []*Task {
	r := make([]*Task, len(l.tasks))
	copy(r, l.tasks)
	return r
}

func (l *List) Clear() {
	l.tasks = nil

	if l.db != nil {
		if _, err := l.db.Exec("drop table if exists tasks"); err != nil {
			log.Error("Couldn't remove table tasks: %s", err)
		}
		l.db.Close()
		l.initDB()
	}
}

func (l *List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Tasks []*Task `json:"tasks"`
	}{
		l.tasks,
	})
}
