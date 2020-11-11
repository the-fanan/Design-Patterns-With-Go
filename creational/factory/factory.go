package factory

import (
	"errors"
	"fmt"
)

type Payment interface {
	Pay(amount float64) string
}

const (
	Cash = 1
	GiftCard = 2
	CreditCard = 3
)

func GetPaymentMethod(m int) (Payment, error) {
	switch m {
		case Cash:
			return new(CashPayment), nil
		case GiftCard:
			return new(GiftCardPayment), nil
		case CreditCard:
			return new(CreditCardPayment), nil
		default:
			return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", m))
	}
}

type CashPayment struct {}

func (c *CashPayment) Pay(amount float64) string {
	return fmt.Sprintf("%f was paid using cash\n", amount)
}


type GiftCardPayment struct {}

func (g *GiftCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("%f was paid using gift card\n", amount)
}

type CreditCardPayment struct {}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("%f was paid using credit card\n", amount)
}