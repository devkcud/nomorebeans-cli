package main

import (
	"fmt"

	"github.com/devkcud/nomorebeans-cli/internal/utils/command"
)

func main() {
	err := command.
		New("nmb").
		WithShortDescription("a cli finance tracking").
		WithAuthors("devkcud").
		WithCommands(
			newRatesCommand(),
			newExchangeCommand(),
		).
		Run()
	if err != nil {
		fmt.Println("Err:", err)
	}
}
