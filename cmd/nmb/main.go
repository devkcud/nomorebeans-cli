package main

import "github.com/devkcud/nomorebeans-cli/internal/utils/command"

func main() {
	command.
		New("nmb").
		WithUsage("a cli finance tracking").
		WithAuthors("devkcud").
		Run()
}
