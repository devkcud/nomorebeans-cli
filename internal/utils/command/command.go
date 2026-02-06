package command

import "github.com/urfave/cli/v3"

type commandWrapper struct {
	cmd *cli.Command
}

func New(name, usage string) *commandWrapper {
	return &commandWrapper{
		cmd: &cli.Command{
			Name:    name,
			Usage:   usage,
			Authors: []any{"devkcud"},
		},
	}
}

func (cw *commandWrapper) WithDescription(description string) *commandWrapper {
	cw.cmd.Description = description
	return cw
}

func (cw *commandWrapper) WithVersion(version string) *commandWrapper {
	cw.cmd.Version = version
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

// TODO: Add WithFlags method
// func (cw *commandWrapper) WithFlags() *commandWrapper
