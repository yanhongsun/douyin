package cache

import (
	"context"
	"sync"
)

// call wraps fn and its results to a struct.
// call 代表一个正在执行或已完成的函数调用。
type call struct {
	// fn is the target function that will be called.
	fn func(ctx context.Context, key interface{}) (interface{}, error)

	// result  err用来记录fn发放执行的返回值
	// 这两个字段在 WaitGroup.Done() 之前被写入（仅写一次），并在 WaitGroup.Done() 之后被读取
	result interface{}
	// err is the failed result of fn.
	err error

	// 用来标识fn方法执行完成之后结果是否立马删除还是保留在singleflight中
	deleted bool

	// 	 用于阻塞调用这个 call 的其他请求
	wg sync.WaitGroup
}

func newCall(fn func(ctx context.Context, key interface{}) (interface{}, error)) *call {
	return &call{
		fn: fn,
	}
}

// 在 Do 函数中，函数先是判断这个 key 是否是第一次调用，如果是，就会进入 doCall 调用回调函数获取结果，
// 后续的请求就会阻塞在 c.wg.Wait() 这里，等待回调函数返回以后，直接拿到结果。
func (c *call) do(ctx context.Context, key interface{}) {
	defer c.wg.Done()
	c.result, c.err = c.fn(ctx, key)
}

// Group stores all calls in flight.
type Group struct {
	// 这里其实没有使用懒加载
	calls map[interface{}]*call

	lock sync.Mutex
}

// NewGroup returns a new Group holder with given mapSize.
func NewGroup(mapSize int) *Group {
	return &Group{
		calls: make(map[interface{}]*call, mapSize),
	}
}

// Call will call fn in single-flight mode.
// 在 Call 函数中，函数先是判断这个 key 是否是第一次调用，如果是，就会进入 doCall 调用回调函数获取结果，
// 后续的请求就会阻塞在 c.wg.Wait() 这里，等待回调函数返回以后，直接拿到结果。
func (g *Group) Call(ctx context.Context, key interface{}, fn func(ctx context.Context, key interface{}) (interface{}, error)) (interface{}, error) {
	g.lock.Lock()
	// 如果 g 中已经存在对该 key 的请求，则不会重复处理 key 的请求 ，所以释放锁，然后阻塞等待已存的 key 请求得到的结果。
	if c, ok := g.calls[key]; ok {
		g.lock.Unlock()
		c.wg.Wait() // Waiting for result...
		return c.result, c.err
	}
	// 如果没有在处理，则创建一个 call ，把 wg 加 1，把 key 存到 calls 中表示已经有在请求了，然后释放锁
	c := newCall(fn)
	c.wg.Add(1)
	g.calls[key] = c
	g.lock.Unlock()

	c.do(ctx, key) // Call fn to get result...

	g.lock.Lock()
	if !c.deleted {
		delete(g.calls, key)
	}
	g.lock.Unlock()
	return c.result, c.err
}

// Delete removes the flight of key so a new flight will start.
func (g *Group) Delete(key interface{}, strstr string) {
	g.lock.Lock()
	defer g.lock.Unlock()

	if c, ok := g.calls[key]; ok {
		c.deleted = true
		delete(g.calls, key)
	}
}

// DeleteAll removes all flights.
func (g *Group) DeleteAll() {
	g.lock.Lock()
	defer g.lock.Unlock()

	for key, c := range g.calls {
		c.deleted = true
		delete(g.calls, key)
	}
}
