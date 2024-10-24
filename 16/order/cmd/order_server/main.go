package main

import (
	"net/http"
	"order/internal/order_server/handler"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	// 商品信息增删改查接口
	// 订单查询接口
	r.GET("/order/:order_id", func(c *gin.Context) {
		orderID := c.Params.ByName("order_id")
		order, product, err := handler.GetOrder(c, orderID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"order": "", "product": ""})
		} else {
			c.JSON(http.StatusOK, gin.H{"order": order, "product": product})
		}
	})

	// 商品信息查询接口
	r.GET("/product/:product_id", func(c *gin.Context) {
		productID := c.Params.ByName("product_id")
		product, err := handler.GetProduct(c, productID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"product": ""})
		} else {
			c.JSON(http.StatusOK, gin.H{"product": product})
		}
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
