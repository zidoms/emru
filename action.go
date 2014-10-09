package emru

import "time"

type action struct {
	before    task
	createdAt time.Time
}
