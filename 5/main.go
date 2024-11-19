package main

import (
	"github.com/bytedance/gopkg/util/gopool"
)

func main() {
	// 创建协程运行
	go func() {
		// do your job
	}()
	// 用默认协程池运行
	gopool.Go(func() {
		/// do your job
	})
}
