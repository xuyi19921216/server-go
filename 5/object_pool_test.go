package main

import (
	"bytes"
	"sync"
	"testing"
)

var data = make([]byte, 1000)

func WriteBufferNoPool() {
	var buf bytes.Buffer
	buf.Write(data)
}
func BenchmarkNoPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 请求处理逻辑
		WriteBufferNoPool()
	}
}

// 对象池
var objectPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

func WriteBufferWithPool() {
	// 获取临时对象
	buf := objectPool.Get().(*bytes.Buffer)
	// 使用
	buf.Write(data)
	// 将buf对象里面的字段恢复初始值
	buf.Reset()
	// 放回
	objectPool.Put(buf)
}
func BenchmarkWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 请求处理逻辑
		WriteBufferWithPool()
	}
}
