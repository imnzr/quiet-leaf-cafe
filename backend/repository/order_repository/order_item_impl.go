package orderrepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/imnzr/quiet-leaf-cafe/backend/helper"
	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

type OrderItemImpl struct {
	DB *sql.DB
}

func NewOrderItems(db *sql.DB) OrderItem {
	return OrderItemImpl{DB: db}
}

// CreateOrderItem implements OrderItem.
func (o OrderItemImpl) CreateOrderItem(ctx context.Context, tx *sql.Tx, request models.OrderRequest) (int64, float64, error) {
	var totalAmount float64
	itemPrices := make([]float64, len(request.Items)) // simpan harga total per item

	// hitung total amount
	for i, items := range request.Items {
		var unitPrice float64
		err := tx.QueryRowContext(ctx, "SELECT price FROM product WHERE product_id = ?", items.Product_id).Scan(&unitPrice)
		if err != nil {
			log.Println("failed to fetch price:", err)
			return 0, 0, err
		}
		itemTotal := unitPrice * float64(items.Quantity)
		itemPrices[i] = itemTotal
		totalAmount += itemTotal
	}

	// order numbers
	orderNumbers := fmt.Sprintf("ORD-%d", time.Now().Unix())
	orderQuery := "INSERT INTO orders(order_number, customer_id, total_amount, status, created_at) VALUES(?,?,?,'pending', NOW())"

	result, err := tx.ExecContext(ctx, orderQuery, orderNumbers, request.Customer_id, totalAmount)
	if err != nil {
		return 0, 0, err
	}

	orderId, err := result.LastInsertId()
	if err != nil {
		log.Println("last insert failed", err)
		return 0, 0, err
	}

	// insert order items
	itemQuery := "INSERT INTO order_items(order_id, product_id, temperature_id, cupsize_id, sweetness_id, topping_id, addon_id, quantity, price) VALUES(?,?,?,?,?,?,?,?,?)"
	for i, item := range request.Items {
		_, err := tx.ExecContext(ctx, itemQuery, orderId, item.Product_id, item.Temperature_id, item.Cupsize_id, item.Sweetness_id, item.Topping_id, item.AddOn_id, item.Quantity, itemPrices[i])
		if err != nil {
			return 0, 0, err
		}
	}
	return orderId, totalAmount, nil
}

// FindOrderById implements OrderItem.
func (o OrderItemImpl) FindOrderById(ctx context.Context, tx *sql.Tx, order_id int64) (models.Order, error) {
	query := "SELECT orders.order_id, orders.order_number, orders.total_amount,orders.status, orders.created_at, customer.name FROM orders JOIN customer ON customer.customer_id = orders.customer_id WHERE orders.order_id = ?"
	rows, err := tx.QueryContext(ctx, query, order_id)
	helper.HandleQueryError(err)

	defer rows.Close()

	order := models.Order{}

	if rows.Next() {
		err := rows.Scan(&order.Order_id, &order.Order_number, &order.TotalAmount, &order.Status, &order.CreatedAt, &order.Customer_name)
		helper.HandleErrorRows(err)
		return order, nil
	} else {
		return models.Order{}, fmt.Errorf("order id with id %d not found", order_id)
	}
}
