package cache

import (
	"time"
)

type value struct {
	// data 真实的value
	data interface{}

	// ttl 过期时间
	ttl time.Duration

	// createTime 创建时间
	createTime time.Time
}

// newValue 创建一个value
func newValue(data interface{}, ttl time.Duration) *value {
	if ttl < 0 {
		ttl = noTTL
	}

	return &value{
		data:       data,
		ttl:        ttl,
		createTime: time.Now(),
	}
}

// alive 判断当前条目是否过期
func (v *value) alive() bool {
	return v != nil && (v.ttl == noTTL || time.Since(v.createTime) <= v.ttl)
}

// renew 重新赋值，这里是为了重用以前申请的空间  在data大小固定的情况下，可使用该函数
func (v *value) renew(data interface{}, ttl time.Duration) *value {
	if v == nil {
		return nil
	}
	// 这里可能存在问题  ，data的大小不一致，会不会导致出错
	v.data = data
	v.ttl = ttl
	v.createTime = time.Now()
	return v
}
