package list

import (
	"database/sql"
	"encoding/json"
	"errors"
	"sort"
	"time"

	log "github.com/limetext/log4go"
	_ "github.com/mattn/go-sqlite3"
	. "github.com/zidoms/emru/list/task"
)

type List struct {
	tasks Tasks
	db    *sql.DB
}

var (
	list *List

	TaskNotFound = errors.New("Task not found")
)

func Emru() *List {
	if list == nil {
		list = newList()
		list.initDB()
		list.load()
	}
	return list
}

func newList() *List {
	l := &List{tasks: make(Tasks, 0)}
	return l
}

func (l *List) initDB() {
	log.Finest("Initializing db")

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

func (l *List) load() {
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
			log.Error(err)
		}
		t := NewTask(title, body)
		t.Id = id
		t.Done = Status(done)
		t.CreatedAt = date
		list.add(t)
	}
	list.flush()
}

func (l *List) flush() {
	sort.Sort(l.tasks)
}

func (l *List) add(t *Task) {
	for l.tasks.Exists(t.Id) {
		t.Id++
	}
	log.Finest("Adding task %v", t)
	l.tasks = append(l.tasks, t)
}

func (l *List) Add(t *Task) Task {
	l.add(t)
	q := "insert into tasks(title, body, done, created_at) values(?, ?, ?, ?)"
	if _, err := l.db.Exec(q, t.Title, t.Body, t.Done.Val(), t.CreatedAt); err != nil {
		log.Error("Couldn't insert task: %s", err)
	}
	return *t
}

func (l *List) remove(i int) {
	if e := len(l.tasks) - 1; i != e {
		copy(l.tasks[i:], l.tasks[i+1:])
	} else {
		l.tasks = l.tasks[:e]
	}
}

func (l *List) Remove(id int) error {
	i := l.tasks.Index(id)
	if i == -1 {
		return TaskNotFound
	}
	l.remove(i)
	if _, err := l.db.Exec("delete from tasks where id = ?", i); err != nil {
		log.Error("Erro on deleting task %d: %s", i, err)
	}
	return nil
}

func (l *List) update(i int, t Task) {
	log.Debug("Updating task %v to %v", l.tasks[i], t)

	nt := l.tasks[i]
	nt.Title = t.Title
	nt.Body = t.Body
	nt.Done = t.Done
}

func (l *List) Update(id int, t Task) error {
	i := l.tasks.Index(id)
	if i == -1 {
		return TaskNotFound
	}
	l.update(i, t)
	q := "update tasks set title = ?, body = ?, done = ? where id = ?"
	if _, err := l.db.Exec(q, t.Title, t.Body, t.Done.Val(), id); err != nil {
		log.Error("Erro on updating task %d: %s", id, err)
	}
	return nil
}

func (l *List) Get(id int) (Task, error) {
	i := l.tasks.Index(id)
	if i == -1 {
		return Task{}, TaskNotFound
	}
	return *l.tasks[i], nil
}

func (l *List) Tasks() Tasks {
	r := make(Tasks, len(l.tasks))
	copy(r, l.tasks)
	return r
}

func (l *List) clear() {
	log.Finest("Clearing all tasks")
	l.tasks = nil
}

func (l *List) Clear() {
	l.clear()

	log.Finest("Removing tasks table from db")
	if _, err := l.db.Exec("drop table if exists tasks"); err != nil {
		log.Error("Couldn't remove table tasks: %s", err)
	}
	l.db.Close()
	l.initDB()
}

func (l *List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Tasks Tasks `json:"tasks"`
	}{
		l.tasks,
	})
}
