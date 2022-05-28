package redisdb

import (
	"errors"
	"strconv"
	"time"
)

func AddCommentIndexCache(videoId, commentNumber int64) error {
	videoIdS := strconv.FormatInt(videoId, 10)
	return RedisClient.Set(videoIdS+"Index", commentNumber, time.Hour*12).Err()
}

func DeleteCommentIndexCache(videoId int64) error {
	videoIdS := strconv.FormatInt(videoId, 10)
	exist, err := CheckCommentIndexCache(videoId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("videoId is not exist in cache")
	}
	return RedisClient.Del(videoIdS + "Index").Err()
}

func UpdateCommentIndexCache(videoId, offset int64) error {
	exist, err := CheckCommentIndexCache(videoId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("videoId is not exist in cache")
	}

	tmp, err := GetCommentIndexCache(videoId)
	if err != nil {
		return err
	}
	tmp += offset
	err = AddCommentIndexCache(videoId, tmp)
	if err != nil {
		return err
	}
	return nil
}

func CheckCommentIndexCache(videoId int64) (bool, error) {
	videoIdS := strconv.FormatInt(videoId, 10)
	exist, err := RedisClient.Exists(videoIdS + "Index").Result()

	if err != nil {
		return false, err
	}

	if exist > 0 {
		return true, nil
	}

	return false, err
}

func GetCommentIndexCache(videoId int64) (int64, error) {
	videoIds := strconv.FormatInt(videoId, 10)
	resS, err := RedisClient.Get(videoIds + "Index").Result()
	if err != nil {
		return 0, err
	}
	res, err := strconv.ParseInt(resS, 10, 64)
	if err != nil {
		return 0, err
	}
	return res, nil
}
