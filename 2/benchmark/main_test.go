package main

import (
	"testing"
	"unsafe"
)

func BenchmarkBytes2StrRaw(b *testing.B) {
	aa := []byte("abcdefg")
	for n := 0; n < b.N; n++ {
		Bytes2StrRaw(aa)
	}
}
func BenchmarkBytes2StrUnsafe(b *testing.B) {
	aa := []byte("abcdefg")
	for n := 0; n < b.N; n++ {
		Bytes2StrUnsafe(aa)
	}
}

func Bytes2StrRaw(b []byte) string {
	return string(b)
}

func Bytes2StrUnsafe(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
