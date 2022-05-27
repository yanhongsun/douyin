package redisdb

import (
	"douyin/kitex_gen/comment"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

func AddCommentsCache(videoId int64, comment []*comment.Comment) error {
	videoIdS := strconv.FormatInt(videoId, 10)
	store, err := json.Marshal(comment)
	if err != nil {
		return err
	}
	return RedisClient.Set(videoIdS, store, time.Hour*12).Err()
}

func DeleteCommentsCache(videoId int64) error {
	videoIdS := strconv.FormatInt(videoId, 10)
	exist, err := CheckCommentsCache(videoId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("videoId is not exist")
	}
	return RedisClient.Del(videoIdS).Err()
}

func CheckCommentsCache(videoId int64) (bool, error) {
	videoIdS := strconv.FormatInt(videoId, 10)
	exist, err := RedisClient.Exists(videoIdS).Result()

	if err != nil {
		return false, err
	}

	if exist > 0 {
		return true, nil
	}

	return false, err
}

func GetCommentsCache(videoId int64) ([]*comment.Comment, error) {
	videoIds := strconv.FormatInt(videoId, 10)
	resS, err := RedisClient.Get(videoIds).Result()
	if err != nil {
		return nil, err
	}
	var res []*comment.Comment
	err = json.Unmarshal([]byte(resS), &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
