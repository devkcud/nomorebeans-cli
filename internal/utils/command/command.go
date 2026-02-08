package command

import (
	"context"
	"os"

	"github.com/urfave/cli/v3"
)

type commandWrapper struct {
	cmd *cli.Command
}

func New(name string) *commandWrapper {
	return &commandWrapper{
		cmd: &cli.Command{
			UseShortOptionHandling: true,
			Suggest:                true,
			Name:                   name,
		},
	}
}

func (cw *commandWrapper) WithUsage(usage string) *commandWrapper {
	cw.cmd.Usage = usage
	return cw
}

func (cw *commandWrapper) WithDescription(description string) *commandWrapper {
	cw.cmd.Description = description
	return cw
}

func (cw *commandWrapper) WithVersion(version string) *commandWrapper {
	cw.cmd.Version = version
	return cw
}

func (cw *commandWrapper) WithAuthors(authors ...string) *commandWrapper {
	for _, author := range authors {
		cw.cmd.Authors = append(cw.cmd.Authors, author)
	}
	return cw
}

func (cw *commandWrapper) WithHookBefore(hook cli.BeforeFunc) *commandWrapper {
	cw.cmd.Before = hook
	return cw
}

func (cw *commandWrapper) WithHookAfter(hook cli.AfterFunc) *commandWrapper {
	cw.cmd.After = hook
	return cw
}

func (cw *commandWrapper) WithAction(action cli.ActionFunc) *commandWrapper {
	cw.cmd.Action = action
	return cw
}

func (cw *commandWrapper) WithCommands(subcommands ...*commandWrapper) *commandWrapper {
	for _, subcmd := range subcommands {
		cw.cmd.Commands = append(cw.cmd.Commands, subcmd.cmd)
	}
	return cw
}

func (cw *commandWrapper) Run() error {
	return cw.cmd.Run(context.Background(), os.Args)
}

func (cw *commandWrapper) WithFlags(flags ...*flagWrapper) *commandWrapper {
	for _, flag := range flags {
		cw.cmd.Flags = append(cw.cmd.Flags, flag.flag)
	}
	return cw
}
