package main

import (
	"sync/atomic"
	"unsafe"
)

// Node 表示栈中的节点
type Node struct {
	next unsafe.Pointer
	v    interface{}
}

// LockFreeStack 表示无锁栈结构体
type LockFreeStack struct {
	top unsafe.Pointer
}

// NewLockFreeStack 创建一个新的无锁栈实例
func NewLockFreeStack() *LockFreeStack {
	return &LockFreeStack{}
}

// Push 将元素压入栈顶
func (s *LockFreeStack) Push(value interface{}) {
	item := Node{v: value}
	for {
		top := atomic.LoadPointer(&s.top)
		item.next = top
		if atomic.CompareAndSwapPointer(&s.top, top, unsafe.Pointer(&item)) {
			return
		}
	}
}

// Pop 从栈顶弹出元素
func (s *LockFreeStack) Pop() (interface{}, bool) {
	for {
		top := atomic.LoadPointer(&s.top)
		if top == nil {
			return nil, false
		}
		item := (*Node)(top)
		next := atomic.LoadPointer(&item.next)
		if atomic.CompareAndSwapPointer(&s.top, top, next) {
			return item.v, true
		}
	}
}
