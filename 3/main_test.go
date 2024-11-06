package main

import (
	"testing"
)

var slices []int

func init() {
	slices = getRawSlices()
}

/*
	 func BenchmarkGetRawSet(b *testing.B) {
		for n := 0; n < b.N; n++ {
			getRawSet(slices)
		}
	}
*/
func BenchmarkGetEmptyStructSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getEmptyStructSet(slices)
	}
}

func BenchmarkGetCapacitySet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getCapacitySet(slices)
	}
}

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
