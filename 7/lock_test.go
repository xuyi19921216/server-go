package main

import (
	"sync"
	"testing"
	"time"
)

const (
	cost     = 1 * time.Millisecond
	readCnt  = 90
	writeCnt = 1
)

// 互斥锁缓存
func BenchmarkMutexReadMore(b *testing.B) {
	c := NewMutexCache() // 互斥锁实现本地缓存
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 1000; j++ {
			for k := 0; k < readCnt; k++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					c.Get("11")
					time.Sleep(cost)

				}()
			}
			for k := 0; k < writeCnt; k++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					c.Set("11", "11")
					time.Sleep(cost)

				}()
			}
		}
		wg.Wait()
	}
}

// 读写锁缓存
func BenchmarkRwReadMore(b *testing.B) {
	c := NewRWMutexCache() // 互斥锁实现本地缓存
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for j := 0; j < 1000; j++ {
			for k := 0; k < readCnt; k++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					c.Get("11")
					time.Sleep(cost)

				}()
			}
			for k := 0; k < writeCnt; k++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					c.Set("11", "11")
					time.Sleep(cost)

				}()
			}
		}
		wg.Wait()
	}
}
