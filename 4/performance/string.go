package performance

import (
	"unsafe"
)

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
