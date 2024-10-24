package handler

import (
	"context"
	"order/internal/pkg/dal/mysql"
	"order/internal/pkg/model"
	"order/internal/pkg/service"
)

// GetOrder 获取订单信息
func GetOrder(ctx context.Context, orderID string) (*model.Order, *model.Product, error) {
	order, err := mysql.GetOrder(ctx, orderID)
	if err != nil {
		return order, nil, nil
	}
	var product *model.Product
	if order != nil {
		// 从service层读商品信息
		product, err = service.GetProduct(ctx, order.ProductID)
	}
	return order, product, err
}
