package orderitemrepository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/imnzr/quiet-leaf-cafe/backend/models"
)

type OrderItemImpl struct{}

func NewOrderItems() OrderItem {
	return OrderItemImpl{}
}

// CreateOrderItem implements OrderItem.
func (o OrderItemImpl) CreateOrderItem(ctx context.Context, tx *sql.Tx, request models.OrderRequest) (int64, error) {
	var totalAmount float64
	itemPrices := make([]float64, len(request.Items)) // simpan harga total per item

	// hitung total amount
	for i, items := range request.Items {
		var unitPrice float64
		err := tx.QueryRowContext(ctx, "SELECT price FROM product WHERE product_id = ?", items.Product_id).Scan(&unitPrice)
		if err != nil {
			log.Println("failed to fetch price:", err)
			return 0, err
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
		return 0, err
	}

	orderId, err := result.LastInsertId()
	if err != nil {
		log.Println("last insert failed", err)
		return 0, err
	}

	// insert order items
	itemQuery := "INSERT INTO order_items(order_id, product_id, temperature_id, cupsize_id, sweetness_id, topping_id, addon_id, quantity, price) VALUES(?,?,?,?,?,?,?,?,?)"
	for i, item := range request.Items {
		_, err := tx.ExecContext(ctx, itemQuery, item.Order_id, item.Product_id, item.Temperature_id, item.Cupsize_id, item.Sweetness_id, item.Topping_id, item.AddOn_id, item.Quantity, itemPrices[i])
		if err != nil {
			return 0, err
		}
	}
	return orderId, nil
}
