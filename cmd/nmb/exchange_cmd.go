package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/devkcud/nomorebeans-cli/internal/utils/command"
	"github.com/devkcud/nomorebeans-cli/internal/utils/currency"
	"github.com/devkcud/nomorebeans-cli/internal/utils/generic"
	"github.com/urfave/cli/v3"
)

func newExchangeCommand() *command.Command {
	return command.
		New("exchange", "ex").
		WithUsage("[--from=rate] --to=rate [--value=float | value]").
		WithFlags(
			command.
				NewFlag[bool]("simple", "s", "minimal", "min").
				WithUsage("Switch to a minimal output").
				WithDefaultValue(false),
			command.
				NewFlag[string]("from").
				WithUsage("Base currency to use in the exchange (check supported `rates` command)").
				WithDefaultValue(currency.StringifyUnsafe(currency.Base)),
			command.
				NewFlag[string]("to").
				WithUsage("Exchange to this currency (check supported `rates` command)"),
			command.
				NewFlag[float64]("value", "v").
				WithUsage("Value to convert").
				WithDefaultValue(1.0),
		).
		WithAction(func(ctx context.Context, c *cli.Command) error {
			if c.String("to") == "" {
				return errors.New("define a --to value")
			}

			from, err := currency.ParseCurrency(c.String("from"))
			if err != nil {
				return err
			}

			to, err := currency.ParseCurrency(c.String("to"))
			if err != nil {
				return err
			}

			value := c.Float64("value")

			if c.Args().Len() > 1 {
				return generic.ErrTooManyArguments
			}
			if c.Args().Len() == 1 {
				if c.IsSet("value") {
					return errors.New("use either --value or positional value, not both")
				}
				value, err = strconv.ParseFloat(c.Args().First(), 64)
				if err != nil {
					return err
				}
			}

			rate, err := currency.Exchange(from, to)
			if err != nil {
				return err
			}

			result := float64(rate) * value

			if c.Bool("simple") {
				fmt.Printf("%.6f\n", result)
				return nil
			}

			fmt.Printf(
				"%.6f %s = %.6f %s\n",
				value,
				strings.ToUpper(currency.StringifyUnsafe(from)),
				result,
				strings.ToUpper(currency.StringifyUnsafe(to)),
			)

			return nil
		})
}
