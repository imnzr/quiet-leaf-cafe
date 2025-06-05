package helper

import (
	"fmt"

	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

func HandleErrorRows(err error) (models.Customer, error) {
	if err != nil {
		return models.Customer{}, fmt.Errorf("failed to scan row: %w", err)
	}
	return models.Customer{}, nil

}
