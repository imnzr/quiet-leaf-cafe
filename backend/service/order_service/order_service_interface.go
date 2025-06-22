package orderservice

import (
	"context"

	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

type OrderService interface {
	CreateOrder(ctx context.Context, request models.OrderRequest) (int64, error)
}
