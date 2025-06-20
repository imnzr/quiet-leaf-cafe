package temperaturerepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnzr/quiet-leaf-cafe/backend/helper"
	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

type TemperatureRepositoryImpl struct{}

// ChoiceTemperature implements TemperatureRepository.
func (t TemperatureRepositoryImpl) ChoiceTemperature(ctx context.Context, tx *sql.Tx) ([]models.Temperature, error) {
	query := "SELECT temperature_id, name FROM `temperature`"
	rows, err := tx.QueryContext(ctx, query)
	helper.HandleQueryError(err)

	defer rows.Close()

	var temperatures []models.Temperature
	for rows.Next() {
		temperature := models.Temperature{}
		err := rows.Scan(&temperature.Temperature_id, &temperature.Name)
		if err != nil {
			fmt.Println("error scanning rows:", err)
			continue
		}
		temperatures = append(temperatures, temperature)
	}
	return temperatures, nil
}

func NewTemperatureRepository() TemperatureRepository {
	return TemperatureRepositoryImpl{}
}
