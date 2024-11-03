package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
)

func main() {
	// cpu采样
	f, err := os.Create("profile.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()
	// 运行业务逻辑
	expensiveCPU()
}

//go:noinline
func expensiveCPU() {
	var sum float32
	for i := 0; i < 10000000; i++ {
		sum += rand.Float32()
	}

	anotherExpensiveCPU()
}

//go:noinline
func anotherExpensiveCPU() {
	var sum int
	for i := 0; i < 1000000; i++ {
		sum += rand.Intn(10)
	}
}
