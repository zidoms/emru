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

func newTask(title, body string) *task {
	t := &task{title, body, false}
	t.actions[0] = NewAction(t)
	t.createdAt = time.Now()
}
