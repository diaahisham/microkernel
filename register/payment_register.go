package register

import (
	"errors"
	"microkernel/interfaces"
)

func PaymentRegister(paymentMethod string) (interfaces.PaymentPluginInterface, error) {
	switch paymentMethod {
	default:
		return nil, errors.New("there are no payment methods registered yet")

	}
}
