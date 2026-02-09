package command

import "github.com/urfave/cli/v3"

type Flag[T any] struct {
	flag cli.Flag
	spec flagSpec[T]
}

type aflag interface {
	cliFlag() cli.Flag
}

func NewFlag[T any](long string, short ...string) *Flag[T] {
	key := *new(T)

	specAny, ok := flagSpecs[key]
	if !ok {
		panic("unsupported flag type")
	}

	spec := specAny.(flagSpec[T])

	return &Flag[T]{
		flag: spec.new(long, short),
		spec: spec,
	}
}

func (f *Flag[T]) WithUsage(usage string) *Flag[T] {
	f.spec.setUsage(f.flag, usage)
	return f
}

func (f *Flag[T]) WithDefaultValue(value T) *Flag[T] {
	f.spec.setDefault(f.flag, value)
	return f
}

func (f *Flag[T]) cliFlag() cli.Flag {
	return f.flag
}
