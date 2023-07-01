package paymentPlugins

import (
	"encoding/json"
	"errors"
	"fmt"
	"microkernel/interfaces"
)

type CardPaymentModel struct {
	CardNumber string  `json:"card_number"`
	Password   string  `json:"password"`
	Amount     float64 `json:"amount"`
}

type CreditCardPayment struct {
	PaymentModel CardPaymentModel
}

func NewCreditCardPayment() interfaces.PaymentPluginInterface {

	return CreditCardPayment{}

}

func (c CreditCardPayment) Pay(payload interface{}) (string, error) {
	payment := CardPaymentModel{}
	jsonData, _ := json.Marshal(payload)
	err := json.Unmarshal(jsonData, &payment)
	if err == nil {
		return fmt.Sprintf("payment done for card number: %s, with amount: %v", payment.CardNumber, payment.Amount), nil
	} else {
		return "", errors.New("wrong payload for credit card payment")
	}
}

func (c CreditCardPayment) Refund(payload interface{}) (string, error) {

	payment := CardPaymentModel{}
	jsonData, _ := json.Marshal(payload)
	err := json.Unmarshal(jsonData, &payment)
	if err == nil {
		return fmt.Sprintf("refund done for card number: %s, with amount: %v", payment.CardNumber, payment.Amount), nil
	} else {
		return "", errors.New("wrong payload for credit card refund")
	}
}
