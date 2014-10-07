package emru

import "time"

type (
	task struct {
		title     string
		body      string
		done      status
		reminder  time.Time
		actions   []*action
		labels    []*label
		createdAt time.Time
	}

	status bool
)
