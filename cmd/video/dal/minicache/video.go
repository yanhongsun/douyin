package minicache

import (
	"context"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/dal/minicache/cache"
	"errors"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

//缓存内容 videos信息 key:value(map[string]=interface{})
//key设计  video:id
// id
// user_id
// play_url
// cover_url
// favorite_count
// comment_count
// title
//create_time

//发布列表 用户发布信息  key:value([]int切片)
//key设计 publish_list:user_id
//videos_id //如果不在缓存则查数据库更新缓存
//用户信息查表
//TODO查询
//查询发布列表：
//命中返回发布列表,user_id->list<vide_id>,在缓存里查询对应videos，若不命中则查数据库
//查询feed流
//查询数据库，返回videos_id列表，在cache查询video，若不存在则查询数据库
//目的：在查询feed流时，减少回表，权衡发布列表与feed流功能之间性能权衡，短视频业务场景：一段时间的热点数据与热点用户往往具有很大交集
//TODO更新
//查询原则只跟cache交互
//上传视频时查询上传用户的发布列表在不在缓存，如果在则将过期时间设为0，，将用户发布列表锁住，当更新数据库后，解锁
//
//
//
//
func GetVideo(ctx context.Context, videoId int64) (db.Video, error) {
	key := cache.StringInt64(videoId)
	// 随机选择1-60的过期时间
	rand.Seed(time.Now().UnixNano())
	sum32 := rand.Intn(60)
	sum64 := int64(sum32)
	sum64 += 1
	onMissed := func(ctx context.Context, key interface{}) (data interface{}, err error) {
		keys := int64(key.(cache.StringInt64))
		return db.GetVideo(ctx, keys)
	}
	tmpStr := "video"
	//fmt.Println("-----------------------------------------------------------------------")
	v, err, _ := Caches.Get(key, tmpStr, cache.WithOpOnMissed(onMissed), cache.WithOpTTL(time.Duration(sum64)*time.Second), cache.WithOpContext(ctx))

	video := (db.Video)(v.(db.Video))

	return video, err
}

func PublishVideo(ctx context.Context, vid db.Video) error {
	//通过DB.WithContext(ctx) 对Context 支持，可为*gorm.DB 设置超时 Context等
	//select()
	if err := db.PublishVideo(ctx, vid); err != nil {
		return err
	} else {
		vid := cache.StringInt64(vid.UserId)
		Caches.DeleteVideoCache(ctx, vid, "user")
	}

	return nil
}

//查询用户发布列表
func GetPublishList(ctx context.Context, userId int64) ([]*db.Video, error) {
	//
	vid := make([]*db.Video, 0, 30)
	key := cache.StringInt64(userId)
	// 随机选择1-60的过期时间
	rand.Seed(time.Now().UnixNano())
	sum32 := rand.Intn(60)
	sum64 := int64(sum32)
	sum64 += 1
	onMissed := func(ctx context.Context, key interface{}) (data interface{}, err error) {
		keys := int64(key.(cache.StringInt64))
		return db.GetPublishListVideoId(ctx, keys)
	}
	tmpStr := "user"
	//fmt.Println("-----------------------------------------------------------------------")
	v, err, _ := Caches.Get(key, tmpStr, cache.WithOpOnMissed(onMissed), cache.WithOpTTL(time.Duration(sum64)*time.Second), cache.WithOpContext(ctx))

	vidList := ([]int64)(v.([]int64))

	for _, id := range vidList {
		video, err := GetVideo(ctx, id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		vid = append(vid, &video)
	}

	return vid, err
}

func GetFeed(ctx context.Context, lastTime int64, limit int) ([]*db.Video, error) {
	var vid []*db.Video
	vidList, err := db.GetFeedVideoId(ctx, lastTime, limit)
	if err != nil {
		return nil, err
	}
	for _, id := range vidList {
		video, err := GetVideo(ctx, id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		vid = append(vid, &video)
	}

	return vid, nil

}

func VerifyVideoId(ctx context.Context, videoId int64) (bool, error) {

	_, err := GetVideo(ctx, videoId)
	if err == nil {
		return true, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return false, err
}
