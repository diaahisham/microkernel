package core

type PaymentRequestModel struct {
	PaymentMethod string      `json:"payment_method"`
	Payload       interface{} `json:"payload"`
}
