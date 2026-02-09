package currency

import (
	"iter"
	"maps"
)

type Currency int

const (
	// Unknown values are just for show purpose. They should not hold meaningful value.
	unknown Currency = -1

	BRL Currency = iota
	USD
	EUR
	CAD
	GBP
	NZD
	PLN
)

// Used as the base currency when implicit by a method
const Base Currency = BRL

var currencyToString = map[Currency]string{
	BRL: "brl",
	USD: "usd",
	EUR: "eur",
	CAD: "cad",
	GBP: "gbp",
	NZD: "nzd",
	PLN: "pln",
}
var stringToCurrency = func() map[string]Currency {
	m := make(map[string]Currency, len(currencyToString))
	for k, v := range currencyToString {
		m[v] = k
	}
	return m
}()

func Supported() iter.Seq[Currency] {
	return maps.Keys(currencyToString)
}
