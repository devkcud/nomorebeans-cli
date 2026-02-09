package main

import "github.com/devkcud/nomorebeans-cli/internal/utils/command"

func main() {
	command.
		New("nmb").
		WithShortDescription("a cli finance tracking").
		WithAuthors("devkcud").
		WithCommands(
			newRatesCommand(),
			newExchangeCommand(),
		).
		Run()
}
