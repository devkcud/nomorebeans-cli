package currency

import "errors"

var (
	ErrUnknownCurrencyId  = errors.New("unknown currency id")
	ErrInvalidAPIResponse = errors.New("invalid API response format; internal method error")
)
