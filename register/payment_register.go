package register

import (
	"errors"
	"microkernel/interfaces"
	"microkernel/paymentPlugins"
)

func PaymentRegister(paymentMethod string) (interfaces.PaymentPluginInterface, error) {
	switch paymentMethod {
	case "creditcard":
		return paymentPlugins.NewCreditCardPayment(), nil
	default:
		return nil, errors.New("there are no payment methods registered yet")

	}
}
