package emru

import "time"

type list struct {
	tasks []*task
	date  time.Time
}
