package model

type Product struct {
	ProductID   string  `json:"product_id"`   // 商品id
	ProductName string  `json:"product_name"` // 商品名称
	Amount      float64 `json:"amount"`       // 商品金额
}
