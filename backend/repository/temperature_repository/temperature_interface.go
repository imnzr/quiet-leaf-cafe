package temperaturerepository

import (
	"context"
	"database/sql"

	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

type TemperatureRepository interface {
	ChoiceTemperature(ctx context.Context, tx *sql.Tx) ([]models.Temperature, error)
}
