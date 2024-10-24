package handler

import (
	"context"
	"order/internal/pkg/model"
	"order/internal/pkg/service"
)

// GetProduct 获取商品信息
func GetProduct(ctx context.Context, productID string) (*model.Product, error) {
	// 从service层读商品信息
	return service.GetProduct(ctx, productID)
}
