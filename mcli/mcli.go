package mcli

import (
	"fmt"
	sound "github.com/noguchidaisuke/cli/mp3-player"
	"github.com/urfave/cli"
	"path/filepath"
)

func New(name, usage, version string) *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Usage = usage
	app.Version = version
	app.Commands = getCommands()
	return app
}

func getCommands()[]cli.Command {
	return []cli.Command {
		{
			Name: "play",
			Usage: "play music",
			Action: play,
		},
	}
}

func play(c *cli.Context) {
	for _, v := range c.Args() {
		_, filename := filepath.Split(v)
		fmt.Printf("start $s...", filename)
		err := sound.Sound(v)
		if err != nil {
			fmt.Printf("cannnot start skipping %s \n", filename)
		}

	}
}