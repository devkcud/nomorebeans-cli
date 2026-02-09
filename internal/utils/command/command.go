package command

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

type Command struct {
	cmd *cli.Command
}

func New(name string, aliases ...string) *Command {
	return &Command{
		cmd: &cli.Command{
			UseShortOptionHandling: true,
			Suggest:                true,
			Name:                   name,
			Aliases:                aliases,
		},
	}
}

func (c *Command) WithUsage(usage string) *Command {
	c.cmd.UsageText = usage
	return c
}

func (c *Command) WithShortDescription(description string) *Command {
	c.cmd.Usage = description
	return c
}

func (c *Command) WithLongDescription(description string) *Command {
	c.cmd.Description = description
	return c
}

func (c *Command) WithVersion(version string) *Command {
	c.cmd.Version = version
	return c
}

func (c *Command) WithAuthors(authors ...string) *Command {
	for _, author := range authors {
		c.cmd.Authors = append(c.cmd.Authors, author)
	}
	return c
}

func (c *Command) WithHookBefore(hook cli.BeforeFunc) *Command {
	c.cmd.Before = hook
	return c
}

func (c *Command) WithHookAfter(hook cli.AfterFunc) *Command {
	c.cmd.After = hook
	return c
}

func (c *Command) WithAction(action cli.ActionFunc) *Command {
	c.cmd.Action = action
	return c
}

func (c *Command) WithCommands(subcommands ...*Command) *Command {
	for _, subcmd := range subcommands {
		c.cmd.Commands = append(c.cmd.Commands, subcmd.cmd)
	}
	return c
}

func (c *Command) Run() error {
	return c.cmd.Run(context.Background(), os.Args)
}

func (c *Command) WithFlags(flags ...aflag) *Command {
	for _, flag := range flags {
		c.cmd.Flags = append(c.cmd.Flags, flag.cliFlag())
	}
	return c
}
