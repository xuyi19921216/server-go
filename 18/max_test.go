package main

import (
	"server-go/18/reflection"
	"server-go/18/regular"
	"testing"
)

// MaxInt函数benchmark
func BenchmarkRegular(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regular.MaxInt(1, 2)
	}
}

// 反射实现的最大值函数benchmark
func BenchmarkReflection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflection.Max(1, 2)
	}
}
