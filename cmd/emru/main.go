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
	app.Usage = "cli interface for emru"
	app.Version = "0.3.0-alpha"
	app.Commands = []cli.Command{
		cmd.Lists,
		cmd.ShowTasks,
		cmd.AddTask,
		cmd.ToggleTask,
		cmd.RemoveTask,
	}

	app.Run(os.Args)
}
