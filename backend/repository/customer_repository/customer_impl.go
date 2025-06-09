package customerrepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnzr/quiet-leaf-cafe/backend/helper"
	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

type CustomerRepositoryImpl struct{}

func NewCustomerRepository() CustomerRepository {
	return CustomerRepositoryImpl{}
}

// FindByEmail implements CustomerRepository.
func (c CustomerRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (models.Customer, error) {
	query := "SELECT customer_id, name, phone_number, email, password FROM `customer` WHERE email = ?"
	rows, err := tx.QueryContext(ctx, query, email)
	helper.HandleQueryError(err)
	defer rows.Close()

	customer := models.Customer{}
	if rows.Next() {
		err := rows.Scan(&customer.Customer_id, &customer.Name, &customer.Phone_number, &customer.Email, &customer.Password)
		helper.HandleErrorRows(err)
		return customer, nil
	} else {
		return models.Customer{}, fmt.Errorf("customer with email %s not found", email)
	}
}

// Delete implements CustomerRepository.
func (c CustomerRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, customer models.Customer) error {
	query := "DELETE FROM customer WHERE customer_id = ?"
	result, err := tx.ExecContext(ctx, query, customer.Customer_id)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %w", err)
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if RowsAffected == 0 {
		return fmt.Errorf("no customer found with id: %d", customer.Customer_id)
	}

	return nil
}

// FindByAll implements CustomerRepository.
func (c CustomerRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) ([]models.Customer, error) {
	query := "SELECT customer_id, name, phone_number, email, password FROM `customer`"
	rows, err := tx.QueryContext(ctx, query)
	helper.HandleQueryError(err)

	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		customer := models.Customer{}
		err := rows.Scan(&customer.Customer_id, &customer.Name, &customer.Phone_number, &customer.Email, &customer.Password)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue // Skip this row if there's an error
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

// FindById implements CustomerRepository.
func (c CustomerRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, customer_id int) (models.Customer, error) {
	query := "SELECT customer_id, name, phone_number, email, password FROM `customer` WHERE customer_id = ?"
	rows, err := tx.QueryContext(ctx, query, customer_id)
	helper.HandleQueryError(err)

	defer rows.Close()

	customer := models.Customer{}

	if rows.Next() {
		err := rows.Scan(&customer.Customer_id, &customer.Name, &customer.Phone_number, &customer.Email, &customer.Password)
		helper.HandleErrorRows(err)
		return customer, nil
	} else {
		return models.Customer{}, fmt.Errorf("customer with id %d not found", customer_id)
	}
}

// Login implements CustomerRepository.
func (c CustomerRepositoryImpl) Login(ctx context.Context, tx *sql.Tx, customer models.Customer) (models.Customer, error) {
	query := "SELECT customer_id, name, phone_number, email FROM `customer` WHERE email = ? AND password = ?"
	rows, err := tx.QueryContext(ctx, query, customer.Email, customer.Password)
	helper.HandleQueryError(err)

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&customer.Customer_id, &customer.Name, &customer.Phone_number, &customer.Email)
		helper.HandleErrorRows(err)
		return customer, nil
	} else {
		return models.Customer{}, fmt.Errorf("invalid email or password")
	}
}

// Save implements CustomerRepository.
func (c CustomerRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, customer models.Customer) (models.Customer, error) {
	query := "INSERT INTO customer (name, phone_number, email, password) VALUES(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, customer.Name, customer.Phone_number, customer.Email, customer.Password)
	helper.HandleQueryError(err)

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		return models.Customer{}, fmt.Errorf("failed to get last insert id: %w", err)
	}
	customer.Customer_id = int(LastInsertId)
	return customer, nil
}

// UpdateName implements CustomerRepository.
func (c CustomerRepositoryImpl) UpdateName(ctx context.Context, tx *sql.Tx, customer models.Customer) (models.Customer, error) {
	query := "UPDATE customer SET name = ? WHERE customer_id = ?"
	result, err := tx.ExecContext(ctx, query, customer.Name, customer.Customer_id)
	helper.HandleQueryError(err)

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Customer{}, fmt.Errorf("failed to get rows affected: %w", err)
	}
	if RowsAffected == 0 {
		return models.Customer{}, fmt.Errorf("no customer found with id: %d", customer.Customer_id)
	}
	return customer, nil
}
