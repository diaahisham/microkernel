package paymentPlugins

import (
	"encoding/json"
	"errors"
	"fmt"
	"microkernel/interfaces"
)

type PaypalPaymentModel struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Amount   float64 `json:"amount"`
}

type PaypalPayment struct {
	PaymentModel PaypalPaymentModel
}

func NewPaypalPayment() interfaces.PaymentPluginInterface {

	return PaypalPayment{}

}

func (c PaypalPayment) Pay(payload interface{}) (string, error) {
	payment := PaypalPaymentModel{}
	jsonData, _ := json.Marshal(payload)
	err := json.Unmarshal(jsonData, &payment)
	if err == nil {
		return fmt.Sprintf("payment done for usernamer: %s, with amount: %v", payment.Username, payment.Amount), nil
	} else {
		return "", errors.New("wrong payload for paypal payment")
	}
}

func (c PaypalPayment) Refund(payload interface{}) (string, error) {

	payment := PaypalPaymentModel{}
	jsonData, _ := json.Marshal(payload)
	err := json.Unmarshal(jsonData, &payment)
	if err == nil {
		return fmt.Sprintf("refund done for usernamer: %s, with amount: %v", payment.Username, payment.Amount), nil
	} else {
		return "", errors.New("wrong payload for paypal refund")
	}
}
