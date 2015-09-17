package cmd

import (
	"strconv"

	"github.com/fatih/color"
	"github.com/zoli/emru/emru"
)

var (
	green = color.New(color.FgGreen).PrintfFunc()
	cyan  = color.New(color.FgCyan).PrintfFunc()
	white = color.New(color.FgWhite).PrintfFunc()
	red   = color.New(color.FgRed).PrintfFunc()

	title = color.New(color.FgMagenta, color.Bold, color.Underline).PrintfFunc()
)

func printLists(ls emru.Lists) {
	white("\n\n")
	printTitles("Name")
	for n, _ := range ls {
		printList(n)
		white("\n")
	}
	white("\n")
	white(strconv.Itoa(len(ls)) + " lists\n")
}

func printList(name string) {
	white(name)
}

func printTasks(ts []emru.Task, list string) {
	white("\n")
	cyan(list + " list")
	white("\n\n")
	printTitles("ID", "Title", "Body", "Done")
	for _, t := range ts {
		printTask(t)
		white("\n")
	}
	white("\n")
}

func printTask(t emru.Task) {
	white(grow(strconv.Itoa(t.ID), size("ID")) + " ")
	white(grow(t.Title, size("Title")) + " ")
	white(grow(t.Body, size("Body")) + " ")
	if t.Done {
		white(grow("✔", size("Done")))
	} else {
		white(grow("✘", size("Done")))
	}
}

func printTitles(ts ...string) {
	for i, t := range ts {
		title(s(t))

		if i < len(ts)-1 {
			space()
		} else {
			white("\n")
		}
	}
}

func s(s string) string {
	return grow(s, size(s))
}

func space() {
	white(" ")
}

func perr(desc string, err error) {
	red("%s: %s\n", desc, err)
}

func size(key string) int {
	switch key {
	case "ID":
		return 3
	case "Name":
		return 16
	case "Title":
		return 24
	case "Body":
		return 40
	case "Done":
		return 6
	}
	return 16
}

func grow(s string, to int) string {
	for len(s) < to {
		s += " "
	}
	return s
}
