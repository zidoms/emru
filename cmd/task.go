package cmd

import (
	"github.com/codegangsta/cli"
)

var (
	Default = func(c *cli.Context) {
		args := c.Args()
		if !args.Present() {
			cli.ShowAppHelp(c)
			return
		}

		if len(args) == 1 {
			ShowTasks(c)
			return
		}

		switch args[1] {
		case "add":
			AddTask.Run(c)
		case "toggle":
			ToggleTask.Run(c)
		case "remove":
			RemoveTask.Run(c)
		default:
			red("unknown command %s\n", c.Args()[1])
		}
	}

	ShowTasks = func(c *cli.Context) {
		list := c.Args().First()
		ts, err := getTasks(list)
		if err != nil {
			perr("Error on getting list tasks", err)
			return
		}
		printTasks(ts, list)
	}

	AddTask = cli.Command{
		Name:  "add",
		Usage: "add new task",
		Action: func(c *cli.Context) {
			list := c.Args().First()
			if err := newTask(list, c.Args()[1]); err != nil {
				perr("Error on creating new task", err)
			}
		},
	}

	ToggleTask = cli.Command{
		Name:  "toggle",
		Usage: "toggle task status",
		Action: func(c *cli.Context) {
			list := c.Args().First()
			if err := toggleTask(list, c.Args()[1]); err != nil {
				perr("Error on toggling task", err)
			}
		},
	}

	RemoveTask = cli.Command{
		Name:  "remove",
		Usage: "remove task",
		Action: func(c *cli.Context) {
			list := c.Args().First()
			if err := deleteTask(list, c.Args()[1]); err != nil {
				perr("Error on delteing task", err)
			}
		},
	}
)
