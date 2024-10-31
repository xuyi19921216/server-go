package mysql

import (
	"context"
	"order/internal/pkg/model"
)

type Product interface {
	GetProduct(ctx context.Context, productID string) (*model.Product, error)
}

type ProductImpl struct {
}

// GetProduct 获取商品信息
func (impl *ProductImpl) GetProduct(ctx context.Context, productID string) (*model.Product, error) {
	var product model.Product
	err := db.QueryRow("SELECT product_id,product_name, amount  FROM orders WHERE product_id =?", productID).Scan(&product.ProductID, &product.ProductName, &product.Amount)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
