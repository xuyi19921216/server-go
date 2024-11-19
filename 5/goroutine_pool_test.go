package main

import (
	"runtime"
	"sync"
	"testing"

	"github.com/bytedance/gopkg/util/gopool"
)

const benchmarkTimes = 10000

func handler(a, b int) int {
	if b < 100 {
		return handler(0, b+1)
	}
	return 0
}

func BenchmarkGoPool(b *testing.B) {
	config := gopool.NewConfig()
	config.ScaleThreshold = 1
	p := gopool.NewPool("benchmark", int32(runtime.GOMAXPROCS(0)), config)
	var wg sync.WaitGroup
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(benchmarkTimes)
		for j := 0; j < benchmarkTimes; j++ {
			p.Go(func() {
				handler(0, 0)
				wg.Done()
			})
		}
		wg.Wait()
	}
}

func BenchmarkGo(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(benchmarkTimes)
		for j := 0; j < benchmarkTimes; j++ {
			go func() {
				handler(0, 0)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkWorkerPool(b *testing.B) {
	workNum := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pool := NewWorkerPool(workNum)
		for j := 0; j < benchmarkTimes; j++ {
			wg.Add(1)
			pool.addTask(func() {
				handler(0, 0)
				wg.Done()
			})

		}
		wg.Wait()
	}

}
