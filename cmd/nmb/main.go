package main

import (
	"context"
	"fmt"
	"sort"

	"github.com/devkcud/nomorebeans-cli/internal/utils/command"
	"github.com/devkcud/nomorebeans-cli/internal/utils/currency"
	"github.com/urfave/cli/v3"
)

func main() {
	command.
		New("nmb").
		WithUsage("a cli finance tracking").
		WithAuthors("devkcud").
		WithCommands(
			command.
				New("rates").
				WithFlags(
					command.
						NewFlag[string]("base", "b").
						WithUsage("Change base rate").
						WithDefaultValue(currency.StringifyUnsafe(currency.Base)),
				).
				WithCommands(
					command.
						New("supported", "sup").
						WithUsage("See supported currencies").
						WithAction(func(ctx context.Context, c *cli.Command) error {
							supported := currency.Supported()

							keys := make([]currency.Currency, 0)
							for k := range supported {
								keys = append(keys, k)
							}

							sort.Slice(keys, func(i, j int) bool {
								return currency.StringifyUnsafe(keys[i]) < currency.StringifyUnsafe(keys[j])
							})

							for _, k := range keys {
								fmt.Printf("%s (%s)\n", currency.StringifyUnsafe(k), currency.StringifyUnsafe(k, currency.StringifyOptions{Friendly: true}))
							}

							fmt.Println()
							fmt.Println("Use the key, not the name (e.g.: brl, usd, cad)")
							return nil
						}),
				).
				WithAction(func(ctx context.Context, c *cli.Command) error {
					base := c.String("base")
					baseCurrency, err := currency.ParseCurrency(base)
					if err != nil {
						return err
					}

					cmap, err := currency.Rates(baseCurrency)
					if err != nil {
						return err
					}

					keys := make([]currency.Currency, 0, len(cmap))
					for k := range cmap {
						keys = append(keys, k)
					}

					sort.Slice(keys, func(i, j int) bool {
						return currency.StringifyUnsafe(keys[i]) < currency.StringifyUnsafe(keys[j])
					})

					fmt.Printf("Base currency: %s\n\n",
						currency.StringifyUnsafe(baseCurrency, currency.StringifyOptions{Friendly: true}),
					)

					fmt.Printf("%-6s %-20s %12s\n", "Code", "Currency", "Rate")
					fmt.Printf("%-6s %-20s %12s\n", "----", "-------------------", "------------")

					for _, k := range keys {
						fmt.Printf(
							"%-6s %-20s %12.6f\n",
							currency.StringifyUnsafe(k),
							currency.StringifyUnsafe(k, currency.StringifyOptions{Friendly: true}),
							cmap[k],
						)
					}

					fmt.Println()
					fmt.Println("Rates are relative to the selected base currency.")

					return nil
				}),
		).
		Run()
}
