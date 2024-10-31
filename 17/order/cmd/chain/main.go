package main

/* import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HelloHandler(w http.ResponseWriter,r *http.Request) {
	// 记录请求开始时间
	start := time.Now()
	defer func() { // 计算响应时间
		elapsed := time.Since(start) // 打印响应时间
		log.Printf("响应时间: %v\n", elapsed)
	}()
	fmt.Fprintf(w, "Hello, World!")
}
func main() {

} */

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 定义一个函数类型，用于表示中间件函数
type Middleware func(http.HandlerFunc) http.HandlerFunc

// 接口耗时记录中间件
func CostMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() { // 计算响应时间
			elapsed := time.Since(start) // 打印响应时间
			log.Printf("响应时间: %v\n", elapsed)
		}()
		next(w, r)
	}
}

// 权限验证中间件
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 这里简单模拟权限验证，假设只有特定的用户可以访问
		if r.Header.Get("Authorization") != "valid_token" {
			http.Error(w, "权限不足", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// 定义实际的业务逻辑处理函数
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

// 应用中间件到处理函数的函数
func ApplyMiddleware(middlewares []Middleware, handler http.HandlerFunc) http.HandlerFunc {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}

func main() {
	// 创建一个HTTP服务器
	http.HandleFunc("/", ApplyMiddleware(
		[]Middleware{CostMiddleware, AuthMiddleware},
		Handler,
	))

	// 启动服务器并监听端口
	log.Fatal(http.ListenAndServe(":8080", nil))
}
