package command

import "github.com/urfave/cli/v3"

type flagSpec[T any] interface {
	new(long string, short []string) cli.Flag
	setUsage(flag cli.Flag, usage string)
	setDefault(flag cli.Flag, value T)
}

type stringFlagSpec struct{}

func (stringFlagSpec) new(long string, short []string) cli.Flag {
	return &cli.StringFlag{Name: long, Aliases: short}
}
func (stringFlagSpec) setUsage(f cli.Flag, usage string) {
	f.(*cli.StringFlag).Usage = usage
}
func (stringFlagSpec) setDefault(f cli.Flag, v string) {
	f.(*cli.StringFlag).Value = v
}

type boolFlagSpec struct{}

func (boolFlagSpec) new(long string, short []string) cli.Flag {
	return &cli.BoolFlag{Name: long, Aliases: short}
}
func (boolFlagSpec) setUsage(f cli.Flag, usage string) {
	f.(*cli.BoolFlag).Usage = usage
}
func (boolFlagSpec) setDefault(f cli.Flag, v bool) {
	f.(*cli.BoolFlag).Value = v
}

type intFlagSpec struct{}

func (intFlagSpec) new(long string, short []string) cli.Flag {
	return &cli.IntFlag{Name: long, Aliases: short}
}
func (intFlagSpec) setUsage(f cli.Flag, usage string) {
	f.(*cli.IntFlag).Usage = usage
}
func (intFlagSpec) setDefault(f cli.Flag, v int) {
	f.(*cli.IntFlag).Value = v
}

type float64FlagSpec struct{}

func (float64FlagSpec) new(long string, short []string) cli.Flag {
	return &cli.Float64Flag{Name: long, Aliases: short}
}
func (float64FlagSpec) setUsage(f cli.Flag, usage string) {
	f.(*cli.Float64Flag).Usage = usage
}
func (float64FlagSpec) setDefault(f cli.Flag, v float64) {
	f.(*cli.Float64Flag).Value = v
}

var flagSpecs = map[any]any{
	*new(string):  stringFlagSpec{},
	*new(bool):    boolFlagSpec{},
	*new(int):     intFlagSpec{},
	*new(float64): float64FlagSpec{},
}
