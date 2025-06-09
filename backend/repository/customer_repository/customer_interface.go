package customerrepository

import (
	"context"
	"database/sql"

	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

type CustomerRepository interface {
	Save(ctx context.Context, tx *sql.Tx, customer models.Customer) (models.Customer, error)
	Delete(ctx context.Context, tx *sql.Tx, customer models.Customer) error
	FindById(ctx context.Context, tx *sql.Tx, customer_id int) (models.Customer, error)
	FindByAll(ctx context.Context, tx *sql.Tx) ([]models.Customer, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (models.Customer, error)
	Login(ctx context.Context, tx *sql.Tx, customer models.Customer) (models.Customer, error)

	UpdateName(ctx context.Context, tx *sql.Tx, customer models.Customer) (models.Customer, error)
}
