package main

import (
	"runtime"
	"testing"
)

// 测试key-value非指针类型,int的gc开销
func Test500wIntGCDuration(t *testing.T) {
	size := 5000000
	m := GenerateIntMap(size)
	runtime.GC()
	gcCost := timeGC()
	t.Logf("size %d GC duration: %v\n", size, gcCost)
	_ = m[1]
}
func GenerateIntMap(size int) map[int]int {
	// 在这里执行一些可能会触发GC的操作，例如创建大量对象等
	// 以下示例创建一个较大的map并填充数据
	m := make(map[int]int)
	for i := 0; i < size; i++ {
		m[i] = i

	}
	return m
}

func TestSmallStruct(t *testing.T) {
	type SmallStruct struct {
		data [128]byte
	}
	m := make(map[int]SmallStruct)
	size := 5000000
	for i := 0; i < size; i++ {
		m[i] = SmallStruct{}
	}
	runtime.GC()
	gcCost := timeGC()
	t.Logf("size %d GC duration: %v\n", size, gcCost)
	_ = m[1]
}
func TestBigStruct(t *testing.T) {
	type BigStruct struct {
		data [129]byte
	}
	m := make(map[int]BigStruct)
	size := 5000000
	for i := 0; i < size; i++ {
		m[i] = BigStruct{}
	}
	runtime.GC()
	gcCost := timeGC()
	t.Logf("size %d GC duration: %v\n", size, gcCost)
	_ = m[1]
}
