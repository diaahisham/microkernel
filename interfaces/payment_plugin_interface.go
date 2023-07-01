package interfaces

type PaymentPluginInterface interface {
	Pay(payload interface{}) (string, error)
	Refund(payload interface{}) (string, error)
}
