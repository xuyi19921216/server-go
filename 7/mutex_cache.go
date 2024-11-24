package main

import (
	"sync"
)

type MutexCache struct {
	mu   sync.Mutex        // 互斥锁
	data map[string]string // 共享数据
}

// NewMutexCache初始化一个MutexCache实例
func NewMutexCache() *MutexCache {
	c := &MutexCache{data: make(map[string]string)}
	return c
}

// Set更新缓存
func (c *MutexCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Get从缓存中获取值
func (c *MutexCache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.data[key]
	return value, ok
}
