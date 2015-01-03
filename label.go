package main

// Task labels
type label struct {
	title string
	color string
}

func newLabel(title, color string) *label {
	return &label{title, color}
}
