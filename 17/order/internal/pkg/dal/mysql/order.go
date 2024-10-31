package mysql

import (
	"context"
	"order/internal/pkg/model"
)

// GetOrder 获取订单信息
func GetOrder(ctx context.Context, orderID string) (*model.Order, error) {
	// 从db获取
	var order model.Order
	err := db.QueryRow("SELECT order_id, amount, product_id FROM orders WHERE order_id =?", orderID).Scan(&order.OrderID, &order.Amount, &order.ProductID)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
