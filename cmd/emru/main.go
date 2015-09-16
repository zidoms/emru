package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/zoli/emru/cmd"
)

func main() {
	app := cli.NewApp()
	app.Email = "zidom72@gmail.com"
	app.Name = "emru"
	// TODO
	// app.Usage = ""
	// app.Version = ""
	app.Commands = []cli.Command{
		cmd.Lists,
		cmd.ShowTasks,
		cmd.AddTask,
		cmd.ToggleTask,
		cmd.RemoveTask,
	}

	app.Run(os.Args)
}
