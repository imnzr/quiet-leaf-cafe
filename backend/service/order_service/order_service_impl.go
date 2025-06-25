package orderservice

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnzr/quiet-leaf-cafe/backend/helper"
	"github.com/imnzr/quiet-leaf-cafe/backend/models"
	paymentservice "github.com/imnzr/quiet-leaf-cafe/backend/payment/payment-service"
	orderrepository "github.com/imnzr/quiet-leaf-cafe/backend/repository/order_repository"
)

type OrderServiceImpl struct {
	OrderRepository orderrepository.OrderItem
	PaymentService  paymentservice.PaymentServerInterface
	DB              *sql.DB
}

// CreateOrder implements OrderService.
func (service *OrderServiceImpl) CreateOrder(ctx context.Context, request models.OrderRequest) (string, error) {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	orderId, totalAmount, err := service.OrderRepository.CreateOrderItem(ctx, tx, request)
	if err != nil {
		return "error", err
	}

	// ambil detail order (email, number, total)
	order, err := service.OrderRepository.FindOrderById(ctx, tx, orderId)
	if err != nil {
		return "error", err
	}

	// buat invoice ke xendit
	paymentUrl, err := service.PaymentService.CreatePayment(orderId, order.Customer_name, order.Order_number, totalAmount)
	if err != nil {
		fmt.Printf("error creating payment: %v\n", err)
	}
	return paymentUrl, nil
}

func NewOrderService(paymentService paymentservice.PaymentServerInterface, orderRepository orderrepository.OrderItem, db *sql.DB) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		PaymentService:  paymentService,
		DB:              db,
	}
}
