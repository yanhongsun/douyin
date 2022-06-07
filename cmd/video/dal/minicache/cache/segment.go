package cache

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// segment 为了并发，所以分段
type segment struct {
	// data key  对应 value地址
	data map[interface{}]*value

	// mapSize map容量
	mapSize int

	// lock
	lock sync.RWMutex
}

// newSegment
func newSegment(mapSize int) *segment {
	return &segment{
		data:    make(map[interface{}]*value, mapSize),
		mapSize: mapSize,
		lock:    sync.RWMutex{},
	}
}

// get
func (s *segment) get(key interface{ ToString(str string) string }, strstr string) (interface{}, bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	keys := key.ToString(strstr)
	if value, ok := s.data[keys]; ok && value.alive() {
		fmt.Println("当前的key:")
		fmt.Println(keys)
		return value.data, true
	}
	return nil, false
}

// set
func (s *segment) set(key interface{ ToString(str string) string }, strstr string, value interface{}, ttl time.Duration) {
	s.lock.Lock()
	defer s.lock.Unlock()
	// 这里可能由于value的大小>v的大小，而导致出错  所以尽量分配一块新的内存
	//if v, ok := s.data[key]; ok {
	//	v.renew(value, ttl) // Reuse value memory
	//	return
	//}
	keys := key.ToString(strstr)
	//fmt.Println("插入的keys")
	//fmt.Println(keys)
	s.data[keys] = newValue(value, ttl)
}

// delete will delete the key in segment.
func (s *segment) delete(ctx context.Context, key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	//TODO 调用数据库更新操作

	// 1. 删除 isfollowed
	// 2. 删除 follow1的关注列表
	//3. 删除 follow2的粉丝列表
	fmt.Println("删除内存中的key:", key)
	delete(s.data, key)
}

// deleteAll 删除所有的data
func (s *segment) deleteAll() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.data = make(map[interface{}]*value, s.mapSize)
}

// size 返回个数
func (s *segment) size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.data)
}

// gc 清理过期的条目
func (s *segment) gc() {
	s.lock.Lock()
	defer s.lock.Unlock()

	for key, value := range s.data {
		if !value.alive() {
			delete(s.data, key)
		}
	}
}
