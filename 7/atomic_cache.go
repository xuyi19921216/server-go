package main

import (
	"sync/atomic"
)

// LockFreeCache结构体用于管理本地缓存
type LockFreeCache struct {
	cacheMap atomic.Value
}

// NewLockFreeCache初始化一个LockFreeCache实例，同时初始化第一个map
func NewLockFreeCache() *LockFreeCache {
	c := &LockFreeCache{}
	initialMap := make(map[string]interface{})
	c.cacheMap.Store(&initialMap)
	return c
}

// Get从缓存中获取值
func (c *LockFreeCache) Get(key string) (interface{}, bool) {
	cacheMap := c.cacheMap.Load().(*map[string]interface{})
	value, ok := (*cacheMap)[key]
	return value, ok
}

// Update替换整个缓存map
func (c *LockFreeCache) Update(newMap map[string]interface{}) {
	newMapPtr := &newMap
	c.cacheMap.Store(newMapPtr)
}
