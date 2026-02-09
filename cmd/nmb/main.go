package main

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/devkcud/nomorebeans-cli/internal/utils/command"
	"github.com/devkcud/nomorebeans-cli/internal/utils/currency"
	"github.com/devkcud/nomorebeans-cli/internal/utils/generic"
	"github.com/urfave/cli/v3"
)

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
