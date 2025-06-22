package orderservice

import (
	"context"
	"database/sql"

	"github.com/imnzr/quiet-leaf-cafe/backend/helper"
	"github.com/imnzr/quiet-leaf-cafe/backend/models"
	orderitemrepository "github.com/imnzr/quiet-leaf-cafe/backend/repository/order_item_repository"
)

type OrderServiceImpl struct {
	OrderRepository orderitemrepository.OrderItem
	DB              *sql.DB
}

// CreateOrder implements OrderService.
func (service *OrderServiceImpl) CreateOrder(ctx context.Context, request models.OrderRequest) (int64, error) {
	tx, err := service.DB.Begin()
	helper.HandleErrorTransaction(err)
	defer helper.HandleTx(tx)

	orderId, err := service.OrderRepository.CreateOrderItem(ctx, tx, request)
	if err != nil {
		return 0, err
	}
	return orderId, nil
}

func NewOrderService(orderRepository orderitemrepository.OrderItem, db *sql.DB) OrderService {
	return &OrderServiceImpl{
		OrderRepository: orderRepository,
		DB:              db,
	}
}
