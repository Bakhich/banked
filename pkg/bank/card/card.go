package card

import (
	"github.com/Bakhich/pkg/bank/types/types.go"
)

func Issue(curreny types.Currency, color string, name string) types.Card {
	return types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: curreny,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}
