package main

import (
	"os"

	"github.com/devkcud/nomorebeans-cli/internal/utils/command"
)

func main() {
	command.
		New(os.Args[0], "a cli finance tracking").
		WithAuthors("devkcud").
		Run()
}
