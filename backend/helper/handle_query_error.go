package helper

import (
	"fmt"

	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

func HandleQueryError(err error) (models.Customer, error) {
	if err != nil {
		return models.Customer{}, fmt.Errorf("failed to execute query: %w", err)
	}
	return models.Customer{}, nil
}
