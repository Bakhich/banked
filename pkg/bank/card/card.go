package card

import (
	"github.com/Bakhich/pkg/bank/types/types.go"
)

func Withdraw(card *types.Card, amount types.Money) {
	const withdrawLimit = 20_000_00
	if amount < 0 {
		return
	}
	if amount > withdrawLimit {
		return
	}
	if !card.Active {
		return
	}
	if card.Balance < amount {
		return
	}

	card.Balance -= amount
}

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
