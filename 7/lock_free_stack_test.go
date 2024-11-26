package main

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewLockFreeStack()

	// 压入元素
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// 弹出元素并打印
	value, ok := stack.Pop()
	if ok {
		fmt.Println("弹出元素:", value)
	}

	value, ok = stack.Pop()
	if ok {
		fmt.Println("弹出元素:", value)
	}

	value, ok = stack.Pop()
	if ok {
		fmt.Println("弹出元素:", value)
	}
}
