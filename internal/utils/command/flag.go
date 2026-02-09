package command

import (
	"reflect"

	"github.com/urfave/cli/v3"
)

type flagWrapper struct {
	flag cli.Flag
}

type flagFactory func(long string, short []string) cli.Flag
type flagKind interface {
	string | bool
}

var factories = map[reflect.Type]flagFactory{
	reflect.TypeFor[string](): func(long string, short []string) cli.Flag {
		return &cli.StringFlag{Name: long, Aliases: short}
	},
	reflect.TypeFor[bool](): func(long string, short []string) cli.Flag {
		return &cli.BoolFlag{Name: long, Aliases: short}
	},
}

func NewFlag[T flagKind](long string, short ...string) *flagWrapper {
	factory, ok := factories[reflect.TypeFor[T]()]
	if !ok {
		panic("unsupported flag type")
	}

	return &flagWrapper{flag: factory(long, short)}
}

func (fw *flagWrapper) WithUsage(usage string) *flagWrapper {
	switch f := fw.flag.(type) {
	case *cli.StringFlag:
		f.Usage = usage
	case *cli.BoolFlag:
		f.Usage = usage
	}

	return fw
}

// WARNING: Make sure that the value is of the same type of the flag
func (fw *flagWrapper) WithDefaultValue(value any) *flagWrapper {
	switch f := fw.flag.(type) {
	case *cli.StringFlag:
		f.Value = value.(string)
	case *cli.BoolFlag:
		f.Value = value.(bool)
	}

	return fw
}
