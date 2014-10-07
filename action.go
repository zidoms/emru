package emru

import "time"

type action struct {
	before    task
	after     task
	createdAt time.Time
}
