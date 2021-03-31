package card

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {

	var choosePayment []types.PaymentSource
	for _, card := range cards {
		if card.Balance < 0 {
			continue
		}
		if !card.Active {
			continue
		}

		choosePayment = append(choosePayment, types.PaymentSource{
			Type:    "card",
			Number:  string(card.PAN),
			Balance: card.Balance,
		})
	}
	return choosePayment
}
