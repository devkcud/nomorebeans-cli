package main

import "github.com/devkcud/nomorebeans-cli/internal/utils/command"

func main() {
	command.NewFlag[string]("output")
}
