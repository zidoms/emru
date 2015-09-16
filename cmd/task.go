package cmd

import (
	"github.com/codegangsta/cli"
)

var (
	ShowTasks = cli.Command{
		Name:  "tasks",
		Usage: "show all list tasks",
		Flags: []cli.Flag{ListFlag},
		Action: func(c *cli.Context) {
			ts, err := getTasks(c.String("list"))
			if err != nil {
				perr("Error on getting list tasks", err)
				return
			}
			printTasks(ts, c.String("list"))
		},
	}

	AddTask = cli.Command{
		Name:  "add",
		Usage: "add new task",
		Flags: []cli.Flag{ListFlag},
		Action: func(c *cli.Context) {
			if err := newTask(c.String("list"), c.Args().First()); err != nil {
				perr("Error on creating new task", err)
			}
		},
	}

	ToggleTask = cli.Command{
		Name:  "toggle",
		Usage: "toggle task status",
		Flags: []cli.Flag{ListFlag},
		Action: func(c *cli.Context) {
			if err := toggleTask(c.String("list"), c.Args().First()); err != nil {
				perr("Error on toggling task", err)
			}
		},
	}

	RemoveTask = cli.Command{
		Name:  "rm",
		Usage: "remove task",
		Flags: []cli.Flag{ListFlag},
		Action: func(c *cli.Context) {
			if err := deleteTask(c.String("list"), c.Args().First()); err != nil {
				perr("Error on delteing task", err)
			}
		},
	}

	ListFlag = cli.StringFlag{
		Name:  "list,l",
		Usage: "set the list you want to take action on",
	}
)
