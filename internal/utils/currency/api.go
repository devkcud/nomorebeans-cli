package currency

import (
	"encoding/json"
	"net/http"
)

// https://github.com/fawazahmed0/exchange-api
// https://github.com/fawazahmed0/exchange-api/issues/90
var apiURLs = []string{
	"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies",
	"https://latest.currency-api.pages.dev/v1/currencies",
	"https://cdn.jsdelivr.net/npm/@fawazahmed0/currency-api@latest/v1/currencies",
	"https://latest.currency-api.pages.dev/v1/currencies",
}

func request(base Currency) (RateMap, error) {
	baseKey, err := Stringify(base)
	if err != nil {
		return nil, err
	}

	var lastErr error

	for _, baseURL := range apiURLs {
		url := baseURL + "/" + baseKey + ".min.json"

		res, err := http.Get(url)
		if err != nil {
			lastErr = err
			continue
		}

		if res.StatusCode != http.StatusOK {
			res.Body.Close()
			lastErr = ErrInvalidAPIResponse
			continue
		}

		var raw map[string]any
		if err := json.NewDecoder(res.Body).Decode(&raw); err != nil {
			res.Body.Close()
			lastErr = err
			continue
		}
		res.Body.Close()

		ratesRaw, ok := raw[baseKey].(map[string]any)
		if !ok {
			return nil, ErrInvalidAPIResponse
		}

		m := make(RateMap, len(currencyToString))
		m[base] = 1

		for c, key := range currencyToString {
			if c == base {
				continue
			}

			v, ok := ratesRaw[key].(float64)
			if !ok || v == 0 {
				continue
			}

			m[c] = Rate(1 / v)
		}

		return m, nil
	}

	return nil, lastErr
}
