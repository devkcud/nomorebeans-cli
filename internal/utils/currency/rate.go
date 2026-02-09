package currency

import "github.com/devkcud/nomorebeans-cli/internal/utils/generic"

type Rate float64

// RateMap[from] = rate to base
//
// 1 unit of `from` equals Rate units of `base`
type RateMap map[Currency]Rate

func Rates(base ...Currency) (RateMap, error) {
	switch len(base) {
	case 0:
		return request(Base)
	case 1:
		return request(base[0])
	default:
		return nil, generic.ErrTooManyArguments
	}
}

func Exchange(from, to Currency) (Rate, error) {
	m, err := Rates(from)
	if err != nil {
		return 0, err
	}

	r, ok := m[to]
	if !ok {
		return 0, ErrUnknownCurrencyId
	}

	return r, nil
}
