package handler

import (
	"context"
	"order/internal/pkg/model"
	"order/internal/pkg/service"
)

// GetProduct 获取商品信息
func GetProduct(ctx context.Context, productID string, scene int) (*model.Product, error) {
	// 调用简单工厂方法创建对象
	impl := service.NewProductService(scene)
	return impl.GetProduct(ctx, productID)
}
