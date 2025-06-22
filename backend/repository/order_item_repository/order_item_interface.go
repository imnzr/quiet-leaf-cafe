package orderitemrepository

import (
	"context"
	"database/sql"

	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

type OrderItem interface {
	CreateOrderItem(ctx context.Context, tx *sql.Tx, request models.OrderRequest) (int64, error)
}
