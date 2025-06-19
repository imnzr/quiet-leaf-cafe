package productrepository

import (
	"context"
	"database/sql"

	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

type ProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, product models.Product) (models.Product, error)
	Delete(ctx context.Context, tx *sql.Tx, product models.Product) error
	Search(ctx context.Context, tx *sql.Tx, keyword string) ([]models.Product, error)
	UpdatePrice(ctx context.Context, tx *sql.Tx, product models.Product) (models.Product, error)
	UpdateDescription(ctx context.Context, tx *sql.Tx, product models.Product) (models.Product, error)
	UpdateName(ctx context.Context, tx *sql.Tx, product models.Product) (models.Product, error)
	FindById(ctx context.Context, tx *sql.Tx, product_id int) (models.Product, error)
	FindByAll(ctx context.Context, tx *sql.Tx) ([]models.Product, error)
}
