package emru

import "time"

type (
	task struct {
		title     string
		body      string
		done      status
		reminder  time.Time
		actions   []*action // Keeps task history
		labels    []*label
		createdAt time.Time
	}

	// Task status is it done or not
	status bool
)

func newTask(title, body string) (t *task) {
	t = &task{title: title, body: body, done: false}
	t.actions[0] = newAction(*t)
	t.createdAt = time.Now()

	return
}

func (s *status) toggle() {
	*s = !(*s)
}
