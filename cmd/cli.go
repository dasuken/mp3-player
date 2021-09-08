package main

import (
	"github.com/noguchidaisuke/cli/mp3-player/mcli"
	"os"
)

func main() {
	app := mcli.New("music-cli", "BGM", "1.0.0")
	app.Run(os.Args)
}