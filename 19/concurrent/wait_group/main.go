package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		/* go func(v int) {
			wg.Add(1) // 错误：在这里调用Add
			// Do some work...
			fmt.Println(v)
			wg.Done()
		}(i) */
		wg.Add(1) // 正确：在协程外部调用Add方法
		go func(v int) {
			defer wg.Done()
			// Do some work...
			fmt.Println(v)
		}(i)
	}
	wg.Wait()
}
