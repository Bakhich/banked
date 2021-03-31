package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSource() {
	cards := []types.Card{
		{
			PAN:     "5058 xxxx xxxx 8888",
			Active:  true,
			Balance: 10_000_00,
		},
		{
			PAN:     "5058 xxxx xxxx 8888",
			Active:  false,
			Balance: 10_000_00,
		},
		{
			PAN:     "5058 xxx xxxx 0000",
			Active:  true,
			Balance: -10_000_00,
		},
	}
	paymentSources := PaymentSources(cards)
	for _, a := range paymentSources {
		fmt.Println(a.Number)
	}

	// Output: 5058 xxxx xxxx 8888
}
