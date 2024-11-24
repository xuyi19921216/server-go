package main

import (
	"sync"
)

// segment结构体代表一个数据分段，包含一个互斥锁和一个map用于存储数据
type segment struct {
	lock sync.RWMutex
	data map[string]string
}

// ConcurrentMap结构体代表整个并发安全的map，包含多个数据分段
type ConcurrentMap struct {
	segments    []*segment
	numSegments int // 定义分段锁的段数，可根据实际情况调整
}

// NewConcurrentMap初始化一个ConcurrentMap实例
func NewConcurrentMap(numSegments int) *ConcurrentMap {
	cm := &ConcurrentMap{
		segments:    make([]*segment, numSegments),
		numSegments: numSegments,
	}
	for i := range cm.segments {
		cm.segments[i] = &segment{
			data: make(map[string]string),
		}
	}
	return cm
}

// getSegmentIndex根据键计算其所属的分段索引
func (cm *ConcurrentMap) getSegmentIndex(key string) int {
	hash := 0
	for _, char := range key {
		hash += int(char)
	}
	return hash % cm.numSegments
}

// Set更新缓存，将键值对存入对应的分段
func (cm *ConcurrentMap) Set(key, value string) {
	segmentIndex := cm.getSegmentIndex(key)
	cm.segments[segmentIndex].lock.Lock()
	defer cm.segments[segmentIndex].lock.Unlock()
	cm.segments[segmentIndex].data[key] = value
}

// Get从缓存中获取值，从对应的分段中查找键对应的值
func (cm *ConcurrentMap) Get(key string) (string, bool) {
	segmentIndex := cm.getSegmentIndex(key)
	cm.segments[segmentIndex].lock.RLock()
	defer cm.segments[segmentIndex].lock.RUnlock()
	value, ok := cm.segments[segmentIndex].data[key]
	return value, ok
}
