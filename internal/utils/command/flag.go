package command

import (
	"reflect"

	"github.com/urfave/cli/v3"
)

type flagWrapper struct {
	flag cli.Flag
}

var flagMap map[string]cli.Flag = map[string]cli.Flag{
	"string": &cli.StringFlag{},
}

func NewFlag[T any](long string, short ...string) *flagWrapper {
	t := reflect.TypeFor[T]()

	flag, exists := flagMap[t.String()]
	if exists {
		return nil
	}

	flag.Set("Name", long)

	return &flagWrapper{flag}
}
