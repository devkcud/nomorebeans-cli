package command

import "github.com/urfave/cli/v3"

type flagWrapper[T any] struct {
	flag cli.Flag
	spec flagSpec[T]
}

func NewFlag[T any](long string, short ...string) *flagWrapper[T] {
	key := *new(T)

	specAny, ok := flagSpecs[key]
	if !ok {
		panic("unsupported flag type")
	}

	spec := specAny.(flagSpec[T])

	return &flagWrapper[T]{
		flag: spec.new(long, short),
		spec: spec,
	}
}

func (fw *flagWrapper[T]) WithUsage(usage string) *flagWrapper[T] {
	fw.spec.setUsage(fw.flag, usage)
	return fw
}

func (fw *flagWrapper[T]) WithDefaultValue(value T) *flagWrapper[T] {
	fw.spec.setDefault(fw.flag, value)
	return fw
}
