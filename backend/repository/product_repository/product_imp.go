package productrepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/imnzr/quiet-leaf-cafe/backend/helper"
	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

type ProductRepositoryImpl struct{}

// FindByAll implements ProductRepository.
func (p ProductRepositoryImpl) FindByAll(ctx context.Context, tx *sql.Tx) ([]models.Product, error) {
	query := "SELECT product_id, name, description, price, image FROM `product`"
	rows, err := tx.QueryContext(ctx, query)
	helper.HandleQueryError(err)

	defer rows.Close()

	var products []models.Product

	for rows.Next() {
		product := models.Product{}
		err := rows.Scan(&product.Product_id, &product.Name, &product.Description, &product.Price, &product.Image)
		if err != nil {
			fmt.Println("error scanning rows:", err)
			continue
		}
		products = append(products, product)
	}
	return products, nil
}

// FindById implements ProductRepository.
func (p ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, product_id int) (models.Product, error) {
	query := "SELECT product_id, name, description, price, image FROM `product` WHERE product_id = ?"
	rows, err := tx.QueryContext(ctx, query, product_id)
	helper.HandleQueryError(err)

	defer rows.Close()

	product := models.Product{}

	if rows.Next() {
		err := rows.Scan(&product.Product_id, &product.Name, &product.Description, &product.Price, &product.Image)
		helper.HandleErrorRows(err)
		return product, nil
	} else {
		return models.Product{}, fmt.Errorf("product with id %d not found", product_id)
	}

}

// Save implements ProductRepository.
func (p ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product models.Product) (models.Product, error) {
	query := "INSERT INTO product(name, description, price, image) VALUES(?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, query, product.Name, product.Description, product.Price, product.Image)
	helper.HandleQueryError(err)

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		return models.Product{}, fmt.Errorf("failed to get last insert id: %w", err)
	}
	product.Product_id = int(LastInsertId)
	return product, nil
}

// Delete implements ProductRepository.
func (p ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product models.Product) error {
	query := "DELETE FROM product WHERE product_id = ?"
	result, err := tx.ExecContext(ctx, query, product.Product_id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if RowsAffected == 0 {
		return fmt.Errorf("no product found with id: %d", product.Product_id)
	}
	return nil
}

// Search implements ProductRepository.
func (p ProductRepositoryImpl) Search(ctx context.Context, tx *sql.Tx, keyword string) ([]models.Product, error) {
	query := "SELECT product_id, name, description, price, image FROM product WHERE name LIKE $1"
	rows, err := tx.QueryContext(ctx, query, "%"+keyword+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Product_id, &product.Name, &product.Price, &product.Image); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

// UpdateDescription implements ProductRepository.
func (p ProductRepositoryImpl) UpdateDescription(ctx context.Context, tx *sql.Tx, product models.Product) (models.Product, error) {
	query := "UPDATE product SET description = ? WHERE product_id = ?"
	result, err := tx.ExecContext(ctx, query, product.Description, product.Product_id)
	helper.HandleQueryError(err)

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Product{}, fmt.Errorf("failed to get rows affected: %w", err)
	}
	if RowsAffected == 0 {
		return models.Product{}, fmt.Errorf("no product found with id: %w", err)
	}
	return product, nil
}

// UpdatePrice implements ProductRepository.
func (p ProductRepositoryImpl) UpdatePrice(ctx context.Context, tx *sql.Tx, product models.Product) (models.Product, error) {
	query := "UPDATE product SET price = ? WHERE product_id = ?"
	result, err := tx.ExecContext(ctx, query, product.Price, product.Product_id)
	helper.HandleQueryError(err)

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Product{}, fmt.Errorf("failed to get rows affected: %w", err)
	}
	if RowsAffected == 0 {
		return models.Product{}, fmt.Errorf("no product found with id: %w", err)
	}
	return product, nil
}

// UpdateTitle implements ProductRepository.
func (p ProductRepositoryImpl) UpdateName(ctx context.Context, tx *sql.Tx, product models.Product) (models.Product, error) {
	query := "UPDATE product SET name = ? WHERE product_id = ?"
	result, err := tx.ExecContext(ctx, query, product.Name, product.Product_id)
	helper.HandleQueryError(err)

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		return models.Product{}, fmt.Errorf("failed to get rows affected: %w", err)
	}
	if RowsAffected == 0 {
		return models.Product{}, fmt.Errorf("no product found with id: %w", err)
	}
	return product, nil
}

func NewProductRepository() ProductRepository {
	return ProductRepositoryImpl{}
}
