package generics

// Max使用泛型来比较两个同类型的值（要求类型是可比较的），并返回较大的值
func Max[T int | float32](a, b T) T {
	if a > b {
		return a
	}
	return b
}
