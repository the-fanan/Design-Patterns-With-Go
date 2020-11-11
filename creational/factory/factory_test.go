package factory

import (
	"strings"
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("A payment method of type 'Cash' must exist")
	}

	msg := payment.Pay(10.3)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message was not correct")
	}
	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodGiftCard(t *testing.T) {
	payment, err := GetPaymentMethod(GiftCard)
	if err != nil {
		t.Fatal("A payment method of type 'GiftCard' must exist")
	}

	msg := payment.Pay(10.3)
	if !strings.Contains(msg, "paid using gift card") {
		t.Error("The gift card payment method message was not correct")
	}
	t.Log("LOG:", msg)
}


func TestCreatePaymentMethodCreditCard(t *testing.T) {
	payment, err := GetPaymentMethod(CreditCard)
	if err != nil {
		t.Fatal("A payment method of type 'CreditCard' must exist")
	}

	msg := payment.Pay(10.3)
	if !strings.Contains(msg, "paid using credit card") {
		t.Error("The credit card payment method message was not correct")
	}
	t.Log("LOG:", msg)
}

func TestCreatePaymentMethodNonExistent(t *testing.T) {
	_, err := GetPaymentMethod(30)
	if err == nil {
		t.Error("A payment method with ID 30 must return an error")
	}
	t.Log("LOG:", err)
}