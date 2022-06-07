package cache

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

var (
	// errNotFound 表示 key 没有找到
	errNotFound = errors.New("cachego: key not found")
)

type StringInt64 int64

func (c StringInt64) ToString(strstr string) string {
	return strconv.FormatInt(int64(c), 10) + strstr
}

type Stringstr string

func (c Stringstr) ToString(strstr string) string {
	return string(c) + strstr
}

type StringInt int

func (c StringInt) ToString(strstr string) string {
	return strconv.Itoa(int(c)) + strstr
}

type StringPairInt64 PairInt64

type PairInt64 struct {
	UserId1 int64
	UserId2 int64
}

func (c StringPairInt64) ToString(strstr string) string {
	return strconv.FormatInt(int64(c.UserId1), 10) + strstr + strconv.FormatInt(int64(c.UserId2), 10)
}

// IsNotFound 返回这个error是不是key没有找到
func IsNotFound(err error) bool {
	return err != nil && err == errNotFound
}

// Cache 是一个cache结构体
type Cache struct {
	// conf 是cache的配置信息
	conf config

	// segments 是存储真实数据的切片，目的是为了提高并发度
	segments []*segment

	// groups 是存储 single_flight keys的切片
	groups []*Group
}

// NewCache : 创建一个新的cache
func NewCache(opts ...Option) *Cache {
	c := &Cache{
		conf: *applyOptions(newDefaultConfig(), opts),
	}

	c.segments = newSegments(c.conf.mapSize, c.conf.segmentSize)
	if c.conf.singleflight {
		c.groups = newGroups(c.conf.mapSize, c.conf.segmentSize)
	}
	// 启动
	// 自动清理过期  每隔一定的时间间隔进行清理
	if c.conf.gcDuration > 0 {
		c.AutoGC(c.conf.gcDuration)
		c.AutoRandGC(c.conf.gcRandDuration)
	}
	return c
}

// newSegments :  初始化每一个分段 segment
func newSegments(mapSize int, segmentSize int) []*segment {
	segments := make([]*segment, segmentSize)
	for i := 0; i < segmentSize; i++ {
		segments[i] = newSegment(mapSize)
	}
	return segments
}

// newGroups : 初始化singleflight groups.
func newGroups(mapSize int, groupSize int) []*Group {
	groups := make([]*Group, groupSize)
	for i := 0; i < groupSize; i++ {
		groups[i] = NewGroup(mapSize)
	}
	return groups
}

// indexOf :对key做一个哈希
func (c *Cache) indexOf(keys interface{ ToString(str string) string }, strstr string) int {

	index := 1469598103934665603
	key := keys.ToString(strstr)
	keyBytes := []byte(key)
	// fmt.Println(key)
	//	time.Sleep(time.Second)
	for _, b := range keyBytes {
		index = (index << 5) - index + int(b&0xff)
		index *= 1099511628211
	}
	return index
}

// segmentOf :根据key对应的哈希求出对应的分段(segment)
func (c *Cache) segmentOf(key interface{ ToString(str string) string }, strstr string) *segment {
	return c.segments[c.indexOf(key, strstr)&(len(c.segments)-1)]
}

// groupOf returns the singleflight group of this key.
func (c *Cache) groupOf(key interface{ ToString(str string) string }, strstr string) *Group {
	return c.groups[c.indexOf(key, strstr)&(len(c.groups)-1)]
}

// Get
// 1. 现在缓存中获取该键的值，如果存在就返回
// 2. 如果没有，则根据指定的函数去获取到该值，并且加载到缓存(可选)
// 3. 如果没有找到，加载到缓存可设置过期时间(可选)
// 4. 如果没有指定函数，则返回IsNotFound
func (c *Cache) Get(key interface{ ToString(str string) string }, strstr string, opts ...OpOption) (interface{}, error, bool) {
	v, ok := c.segmentOf(key, strstr).get(key, strstr)
	if ok {
		fmt.Println("找到cache:")
		fmt.Println(v)
		return v, nil, ok
	}
	fmt.Println("进入cache但是没命中:")
	fmt.Println(key, strstr)
	if len(opts) <= 0 {
		return nil, errNotFound, false
	}
	// TODO
	conf := applyOpOptions(newDefaultGetConfig(), opts)
	if conf.onMissed == nil {
		return nil, errNotFound, false
	}

	var data interface{}
	var err error
	if c.conf.singleflight && conf.singleflight {
		data, err = c.groupOf(key, strstr).Call(conf.ctx, key, conf.onMissed)
	} else {
		data, err = conf.onMissed(conf.ctx, key)
	}

	if err != nil {
		return nil, err, false
	}

	if conf.reload {
		fmt.Println("保存cache")
		c.Set(key, strstr, data, WithOpTTL(conf.ttl))
	}
	return data, nil, true
}

// Set
// 默认情况下，设置的缓存不会过期  可选  过期时间
func (c *Cache) Set(key interface{ ToString(str string) string }, strstr string, value interface{}, opts ...OpOption) {
	conf := applyOpOptions(newDefaultSetConfig(), opts)
	c.segmentOf(key, strstr).set(key, strstr, value, conf.ttl)
}

// DeleteAll
// DeleteAll删除缓存中的所有键。弱一致性
func (c *Cache) DeleteAll() {
	for _, segment := range c.segments {
		segment.deleteAll()
	}

	for _, group := range c.groups {
		group.DeleteAll()
	}
}
func (c *Cache) DeleteVideoCache(ctx context.Context, key interface{ ToString(str string) string }, strstr string) error {
	videoId := StringInt64(key.(StringInt64))
	keytmp := videoId.ToString(strstr)
	c.segmentOf(videoId, strstr).delete(ctx, keytmp)
	return nil
}

// Size 返回cache的大小
// 弱一致性
func (c *Cache) Size() int {
	size := 0

	for _, segment := range c.segments {
		size += segment.size()
	}
	return size
}

// GC 删除了缓存中的过期条目。
// 请注意，此方法是弱一致性的，并且不能保证100％被删除。
func (c *Cache) GC() {

	for _, segment := range c.segments {
		segment.gc()
	}
}
func (c *Cache) RandGc() {
	//fmt.Println("随机RandGc")
	len := len(c.segments)
	rand.Seed(time.Now().UnixNano())
	sum := rand.Intn(len / GcLen)
	for i := 0; i < sum; i++ {
		c.segments[rand.Intn(len)].gc()
	}
}

// AutoGC
// 哨兵模式 清理机制，每隔一定的时间间隔进行清理
func (c *Cache) AutoGC(duration time.Duration) chan<- struct{} {
	quitChan := make(chan struct{})

	go func() {
		ticker := time.NewTicker(duration)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				c.GC()
			case <-quitChan:
				return
			}
		}
	}()

	return quitChan
}

func (c *Cache) AutoRandGC(duration time.Duration) chan<- struct{} {
	quitChan := make(chan struct{})

	go func() {
		ticker := time.NewTicker(duration)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				c.RandGc()
			case <-quitChan:
				return
			}
		}
	}()

	return quitChan
}
