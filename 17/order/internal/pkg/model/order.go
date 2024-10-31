package model

type Order struct {
	OrderID   string  `json:"order_id"`
	Amount    float64 `json:"amount"`     // 订单金额
	ProductID string  `json:"product_id"` // 商品id
}
