package reflection

import "testing"

func TestMax(t *testing.T) {
	a := "aaa"
	b := "bbb"
	_, err := Max(a, b)
	if err != nil {
		panic(err)
	}
}
