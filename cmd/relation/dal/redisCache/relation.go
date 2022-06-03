package redisCache

import (
	"context"
	"douyin/cmd/relation/dal/db"
	"douyin/cmd/relation/dal/redisCache/cache"
	"math/rand"
	"time"
)

func GetFans(ctx context.Context, follower1 int64) ([]db.UserList, error) {
	// 判断是否命中，命中则直接返回；
	// 没有命中，则返回从数据库查询，并更新到缓存（保证缓存是最新的）
	// 实现了singleflight , 多个相同的请求只发一个请求给数据库
	key := cache.StringInt64(follower1)
	// 随机选择1-60的过期时间
	rand.Seed(time.Now().UnixNano())
	sum32 := rand.Intn(60)
	sum64 := int64(sum32)
	sum64 += 1
	onMissed := func(ctx context.Context, key interface{}) (data interface{}, err error) {
		keys := int64(key.(cache.StringInt64))
		return db.GetFans(ctx, keys)
	}
	tmpStr := "Fans"
	//fmt.Println("-----------------------------------------------------------------------")
	v, err, _ := Caches.Get(key, tmpStr, cache.WithOpOnMissed(onMissed), cache.WithOpTTL(time.Duration(sum64)*time.Second), cache.WithOpContext(ctx))

	user_list := ([]db.UserList)(v.([]db.UserList))

	return user_list, err
}
func GetFollows(ctx context.Context, follower1 int64) ([]db.UserList, error) {
	// 判断是否命中，命中则直接返回；
	// 没有命中，则返回从数据库查询，并更新到缓存（保证缓存是最新的）
	// 实现了singleflight , 多个相同的请求只发一个请求给数据库
	key := cache.StringInt64(follower1)
	// 随机选择1-60的过期时间
	rand.Seed(time.Now().UnixNano())
	sum32 := rand.Intn(60)
	sum64 := int64(sum32)
	sum64 += 1
	onMissed := func(ctx context.Context, key interface{}) (data interface{}, err error) {
		keys := int64(key.(cache.StringInt64))
		return db.GetFollows(ctx, keys)
	}
	tmpStr := "Follows"
	// fmt.Println("-----------------------------------------------------------------------")
	// 后面加上ctx的op操作
	v, err, _ := Caches.Get(key, tmpStr, cache.WithOpOnMissed(onMissed), cache.WithOpTTL(time.Duration(sum64)*time.Second), cache.WithOpContext(ctx))

	user_list := ([]db.UserList)(v.([]db.UserList))

	return user_list, err
}
func IsFollowed(ctx context.Context, userId, otherId int64) (bool, error) {
	tmp := cache.PairInt64{
		UserId1: userId,
		UserId2: otherId,
	}
	key := cache.StringPairInt64(tmp)
	rand.Seed(time.Now().UnixNano())
	// 随机选择1-60的过期时间
	sum32 := rand.Intn(60)
	sum64 := int64(sum32)
	sum64 += 1
	onMissed := func(ctx context.Context, key interface{}) (data interface{}, err error) {
		keys := cache.PairInt64(key.(cache.StringPairInt64))
		return db.IsFollowed(ctx, keys.UserId1, keys.UserId2)
	}
	tmpStr := "IsFollowed"
	// fmt.Println("-----------------------------------------------------------------------")
	v, err, _ := Caches.Get(key, tmpStr, cache.WithOpOnMissed(onMissed), cache.WithOpTTL(time.Duration(sum64)*time.Second), cache.WithOpContext(ctx))
	tag := (bool)(v.(bool))
	return tag, err
}
func Follow(ctx context.Context, userId, otherId int64) error {
	tmp := cache.PairInt64{
		UserId1: userId,
		UserId2: otherId,
	}
	key := cache.StringPairInt64(tmp)
	var strstr string
	err := Caches.DeleteFollow(ctx, key, strstr)
	// 删除缓存
	return err
}
func UnFollow(ctx context.Context, userId, otherId int64) error {
	tmp := cache.PairInt64{
		UserId1: userId,
		UserId2: otherId,
	}
	key := cache.StringPairInt64(tmp)
	var strstr string
	err := Caches.DeleteUnFollow(ctx, key, strstr)
	return err
	// 删除缓存
}
