package cmd

import (
	"github.com/codegangsta/cli"
)

var (
	ShowLists = cli.Command{
		Name:  "lists",
		Usage: "show all todo lists",
		Action: func(c *cli.Context) {
			ls, err := getLists()
			if err != nil {
				perr("Error on getting lists", err)
				return
			}

			printLists(ls)
		},
	}

	AddList = cli.Command{
		Name:  "add",
		Usage: "add new list",
		Action: func(c *cli.Context) {
			name := c.Args().First()
			if name == "" {
				return
			}

			if err := newList(name); err != nil {
				perr("Error on creating new list", err)
			}

			white("\n\nCreated new " + name + " list successfully\n")
		},
	}

	RemoveList = cli.Command{
		Name:  "remove",
		Usage: "remove list",
		Action: func(c *cli.Context) {
			name := c.Args().First()
			if name == "" {
				return
			}

			if err := deleteList(name); err != nil {
				perr("Error on removing list", err)
			}

			white("\n\n" + name + " list removed successfully\n")
		},
	}
)
