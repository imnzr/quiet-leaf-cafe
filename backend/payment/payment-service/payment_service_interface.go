package paymentservice

type PaymentServerInterface interface {
	CreatePayment(orderId int64, customerName string, orderNumber string, totalAmount float64) (string, error)
}
