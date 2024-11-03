package main

import (
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("heap.pprof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	runtime.GC() // invoke gc, in order to get up-to-date statistics
	// 业务逻辑
	expensiveMem()

	if err := pprof.WriteHeapProfile(f); err != nil {
		panic(err)
	}
}

//go:noinline
func expensiveMem() {
	m := make([]int, 10000000)
	for i := range m {
		m[i] = rand.Intn(127)
	}

	anotherExpensiveMem()
}

//go:noinline
func anotherExpensiveMem() {
	m := make(map[int]float32, 1000000)

	for key := range m {
		m[key] = rand.Float32()
	}
}
