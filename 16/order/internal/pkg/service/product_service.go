package service

import (
	"context"
	"order/internal/pkg/dal/mysql"
	"order/internal/pkg/dal/redis"
	"order/internal/pkg/model"
)

// GetProduct 获取商品信息
func GetProduct(ctx context.Context, productID string) (*model.Product, error) {
	product, err := redis.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}
	if product != nil {
		return product, nil
	}
	product, err = mysql.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}
	redis.SetProduct(ctx, product)
	return product, nil
}
