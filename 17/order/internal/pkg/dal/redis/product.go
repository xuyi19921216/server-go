package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"order/internal/pkg/model"

	"github.com/go-redis/redis"
)

type Product interface {
	GetProduct(ctx context.Context, productID string) (*model.Product, error)
	SetProduct(ctx context.Context, product *model.Product) error
}

type ProductImpl struct {
}

// GetProduct 获取商品信息
func (impl *ProductImpl) GetProduct(ctx context.Context, productID string) (*model.Product, error) {
	// 从 Redis 中获取商品信息
	productKey := fmt.Sprintf("product:%s", productID)
	result, err := client.Get(productKey).Result()

	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	// 假设商品信息在缓存中是以 JSON 格式存储的
	var product model.Product
	err = json.Unmarshal([]byte(result), &product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (impl *ProductImpl) SetProduct(ctx context.Context, product *model.Product) error {
	return nil
}
