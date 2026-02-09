package currency

import "github.com/devkcud/nomorebeans-cli/internal/utils/generic"

type StringifyOptions struct {
	Friendly bool
}

var friendlyName = map[Currency]string{
	BRL: "Brazilian Real",
	USD: "American Dollar",
	EUR: "Euro",
	CAD: "Canadian Dollar",
	GBP: "British Pound",
	NZD: "New Zeland Dollar",
	PLN: "Polish Zloty",
}

func Stringify(c Currency, opts ...StringifyOptions) (string, error) {
	if len(opts) > 1 {
		return "", generic.ErrTooManyArguments
	}

	m := currencyToString
	if len(opts) == 1 && opts[0].Friendly == true {
		m = friendlyName
	}

	if currency, ok := m[c]; ok {
		return currency, nil
	}

	return "", ErrUnknownCurrencyId
}

// Use only if certain that the id is valid, otherwise it will panic.
// If uncertain, use Stringify.
func StringifyUnsafe(c Currency, opts ...StringifyOptions) string {
	currency, err := Stringify(c, opts...)
	if err != nil {
		panic(err)
	}
	return currency
}

func ParseCurrency(s string) (Currency, error) {
	if currency, ok := stringToCurrency[s]; ok {
		return currency, nil
	}
	return unknown, ErrUnknownCurrencyId
}

// Use only if certain that the id is valid, otherwise it will panic.
// Preferred to use, for example, `currency.BRL` or `currency.Base` instead.
func ParseCurrencyUnsafe(s string) Currency {
	currency, err := ParseCurrency(s)
	if err != nil {
		panic(err)
	}
	return currency
}
