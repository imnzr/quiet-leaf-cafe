package invoiceorders

import (
	orderitemrepository "github.com/imnzr/quiet-leaf-cafe/backend/repository/order_item_repository"
	"github.com/xendit/xendit-go/v7"
)

type PaymentService interface {
	CreatePayment(orderId int, customerEmail string, orderNumber string, totalAmount float64) (string, error)
}

type PaymentServerImpl struct {
	client      *xendit.APIClient
	PaymentRepo orderitemrepository.OrderItem
	OrderRepo
}

func NewPaymentService() PaymentService {

}

// func CreateInvoiceOrders(orderId int, customerEmail string, orderNumber string, totalAmount float64) (string, error) {

// 	// Set secret key dari environment variable untuk keamanan
// 	key := os.Getenv("xendit_key")
// 	if key == "" {
// 		return "", errors.New("xendit secret key is not set in environment")
// 	}
// 	// client api key
// 	xenditClient := xendit.NewClient(key)

// }
