package main

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestAddWithTestify(t *testing.T) {
	a := 3
	b := 5
	result := Add(a, b)
	expected := 8

	// Assert
	assert.Equal(t, expected, result, "Add(3, 5) should return 8")
}

func TestAddWithConvey(t *testing.T) {
	Convey("关于Add函数的测试", t, func() {
		Convey("正常情况的测试", func() {
			Convey("两个正数相加", func() {
				result := Add(2, 3)
				So(result, ShouldEqual, 5)
			})
			Convey("一个正数和一个负数相加", func() {
				result := Add(5, -3)
				So(result, ShouldEqual, 2)
			})
		})
		Convey("边界情况的测试", func() {
			Convey("两个零相加", func() {
				result := Add(0, 0)
				So(result, ShouldEqual, 0)
			})
			Convey("一个数与最大整数相加", func() {
				result := Add(int(math.MaxInt32), 1)
				So(result, ShouldEqual, int(math.MaxInt32)+1)
			})
		})
	})
}

// 测试函数
func TestNonTableDriven(t *testing.T) {
	// 测试用例1
	assert := assert.New(t)
	result := Add(1, 2)
	assert.Equal(result, 3, "add(1, 2) should return 3")

	// 测试用例2
	result = Add(-1, 1)
	assert.Equal(result, 0, "add(-1, 1) should return 0")

	// 测试用例3
	result = Add(0, 0)
	assert.Equal(result, 0, "add(0, 0) should return 0")
}

func TestTableDriven(t *testing.T) {
	tests := []struct {
		name     string
		inputA   int
		inputB   int
		expected int
	}{
		{"Add positive numbers", 1, 2, 3},
		{"Add negative numbers", -1, 1, 0},
		{"Add zero", 0, 0, 0},
		// 更多测试用例...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Add(tt.inputA, tt.inputB)
			assert.Equal(t, tt.expected, result, "结果不符合期望")
		})
	}
}
