package main

import (
	"github.com/codegangsta/cli"
	"github.com/piokaczm/godzira/commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "godzira"
	app.Version = "1.0.3"
	// add more precise description, add some better help text for deploy
	app.Usage = "Smash your apps to servers just like Godzira would smash a city!"
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "build config directory and config file",
			Action: commands.Config,
		},
		{
			Name:   "deploy",
			Usage:  "build and deploy binary to remote server",
			Action: commands.Deploy,
		},
	}

	app.Run(os.Args)
}