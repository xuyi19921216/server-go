package service

import (
	"context"
	"order/internal/pkg/dal/mysql"
	"order/internal/pkg/dal/redis"
	"order/internal/pkg/model"
)

type ProductService interface {
	GetProduct(ctx context.Context, productID string) (*model.Product, error)
}
type ProductServiceCacheImpl struct {
	cache redis.Product
	db    mysql.Product
}

// GetProduct 获取商品信息
func (impl *ProductServiceCacheImpl) GetProduct(ctx context.Context, productID string) (*model.Product, error) {
	product, err := impl.cache.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}
	if product != nil {
		return product, nil
	}
	product, err = impl.db.GetProduct(ctx, productID)
	if err != nil {
		return nil, err
	}
	impl.cache.SetProduct(ctx, product)
	return product, nil
}

func NewProductCacheService(cacheProduct redis.Product, dbProduct mysql.Product) ProductService {
	return &ProductServiceCacheImpl{cache: cacheProduct, db: dbProduct}
}

type ProductServiceDBImpl struct {
	db mysql.Product
}

// GetProduct 获取商品信息
func (impl *ProductServiceDBImpl) GetProduct(ctx context.Context, productID string) (*model.Product, error) {
	return impl.db.GetProduct(ctx, productID)
}

func NewProductDBService(dbProduct mysql.Product) ProductService {
	return &ProductServiceDBImpl{db: dbProduct}
}

// 简单工厂方法，根据传入参数创建同类对象
func NewProductService(scene int) ProductService {
	var impl ProductService
	if scene == 1 {
		// 直接读db
		impl = NewProductDBService(&mysql.ProductImpl{})
	} else {
		// 缓存读写逻辑
		impl = NewProductCacheService(&redis.ProductImpl{}, &mysql.ProductImpl{})
	}
	return impl
}
