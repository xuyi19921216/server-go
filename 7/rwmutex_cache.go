package main

import (
	"sync"
)

type RWMutexCache struct {
	rw   sync.RWMutex      // 读写锁
	data map[string]string // 共享数据
}

// NewRWMutexCache初始化一个RWMutexCache实例
func NewRWMutexCache() *RWMutexCache {
	c := &RWMutexCache{data: make(map[string]string)}
	return c
}

// Set更新缓存
func (c *RWMutexCache) Set(key, value string) {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.data[key] = value
}

// Get从缓存中获取值
func (c *RWMutexCache) Get(key string) (string, bool) {
	c.rw.RLock()
	defer c.rw.RUnlock()
	value, ok := c.data[key]
	return value, ok
}
