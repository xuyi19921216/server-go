package main

import (
	"context"
	"fmt"
)

func main() {
	// 假设有一个大的字符串值
	largeValue := []byte("your_large_value")

	// 指定每个chunk的大小，例如1024字节
	chunkSize := 1024
	ctx := context.TODO()
	// 将大Value存入Redis
	err := storeValueInRedis(ctx, "key", largeValue, chunkSize)
	if err != nil {
		fmt.Println("Error storing large value in Redis:", err)
		return
	}
	data, err := getDataFromRedis(ctx, "key")
	if err != nil {
		fmt.Println("Error get value in Redis:", err)
		return
	}
	fmt.Println(string(data))
}
