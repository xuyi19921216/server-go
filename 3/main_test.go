package main

import (
	"testing"
)

/* var slices []int

func init() {
	slices = getRawSlices()
} */

/*
	 func BenchmarkGetRawSet(b *testing.B) {
		for n := 0; n < b.N; n++ {
			getRawSet(slices)
		}
	}
*/
/* func BenchmarkGetEmptyStructSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getEmptyStructSet(slices)
	}
}

func BenchmarkGetCapacitySet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getCapacitySet(slices)
	}
} */

/* func BenchmarkGetRawSlices(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getRawSlices()
	}
}

func BenchmarkGetCapacitySlices(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getCapacitySlices()
	}
} */

type Item struct {
	id  int
	val [2048]byte
}

// 用下标遍历[]struct{}
func BenchmarkSliceStructByFor(b *testing.B) {
	var items [2048]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for j := 0; j < len(items); j++ {
			tmp = items[j].id
		}
		_ = tmp
	}
}

// 用range的下标遍历[]struct{}
func BenchmarkSliceStructByRangeIndex(b *testing.B) {
	var items [2048]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := range items {
			tmp = items[k].id
		}
		_ = tmp
	}
}

// 用range遍历[]struct{}的元素
func BenchmarkSliceStructByRangeValue(b *testing.B) {
	var items [2048]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
