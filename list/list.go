package list

import (
	"encoding/json"
	"errors"
	"sort"
	"time"

	log "github.com/limetext/log4go"
	. "github.com/zoli/emru/list/task"
)

type (
	List struct {
		tasks     Tasks
		createdAt time.Time
	}

	Lists map[string]*List
)

var TaskNotFound = errors.New("Task not found")

func newList() *List {
	return &List{tasks: make(Tasks, 0), createdAt: time.Now()}
}

func (l *List) flush() {
	sort.Sort(l.tasks)
}

func (l *List) add(t *Task) {
	for l.tasks.Exists(t.Id) {
		t.Id++
	}
	log.Finest("Adding task %v", *t)
	l.tasks = append(l.tasks, t)
}

func (l *List) Add(t *Task) Task {
	l.add(t)
	return *t
}

func (l *List) remove(i int) {
	log.Debug("Removing task %d: %v", i, *l.tasks[i])
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
	return nil
}

func (l *List) update(i int, t Task) {
	log.Debug("Updating task %v to %v", *l.tasks[i], t)
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
	return nil
}

func (l *List) Get(id int) (Task, error) {
	i := l.tasks.Index(id)
	if i == -1 {
		return Task{}, TaskNotFound
	}
	return *l.tasks[i], nil
}

func (l *List) Tasks() []Task {
	r := make([]Task, 0, len(l.tasks))
	for _, t := range l.tasks {
		r = append(r, *t)
	}
	return r
}

func (l *List) clear() {
	log.Finest("Clearing all tasks")
	l.tasks = nil
}

func (l *List) Clear() {
	l.clear()
}

func (l *List) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Tasks     []Task    `json:"tasks"`
		CreatedAt time.Time `json:"created_at"`
	}{
		l.Tasks(),
		l.createdAt,
	})
}
