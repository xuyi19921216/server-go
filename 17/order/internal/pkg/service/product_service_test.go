package service

import (
	"context"
	"order/internal/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProduct(t *testing.T) {
	ctx := context.TODO()

	impl := NewProductCacheService(&MockProductCache{}, &MockProductDB{})
	productID := "1"
	product, _ := impl.GetProduct(ctx, productID)
	assert.Equal(t, productID, product.ProductID, "The two id should be the same.")
}

// mock 商品读写redis
type MockProductCache struct {
}

func (cache *MockProductCache) GetProduct(ctx context.Context, productID string) (*model.Product, error) {
	return &model.Product{ProductID: productID}, nil
}

func (cache *MockProductCache) SetProduct(ctx context.Context, product *model.Product) error {
	return nil
}

// mock商品从db读
type MockProductDB struct {
}

func (db *MockProductDB) GetProduct(ctx context.Context, productID string) (*model.Product, error) {
	return nil, nil
}
