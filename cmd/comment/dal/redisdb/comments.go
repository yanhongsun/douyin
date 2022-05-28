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

func DeleteCommentsCache(videoId, commentId int64) error {
	pipe := RedisClient.TxPipeline()

	// comment_cache change
	exist, err := CheckCommentsCache(videoId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("videoId is not exist in cache")
	}

	tmp, err := GetCommentsCache(videoId)
	if err != nil {
		return err
	}

	index := 0
	for k, value := range tmp {
		if value.CommentId == commentId {
			index = k
			break
		}
	}

	tmp = append(tmp[:index], tmp[index+1:]...)

	videoIdS := strconv.FormatInt(videoId, 10)
	store, err := json.Marshal(tmp)
	if err != nil {
		return err
	}

	if err := pipe.Set(videoIdS, store, time.Hour*12).Err(); err != nil {
		return nil
	}

	// comment_index_cache change
	exist, err = CheckCommentIndexCache(videoId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("videoId is not exist in cache")
	}

	tmpN, err := GetCommentIndexCache(videoId)
	if err != nil {
		return err
	}
	tmpN -= 1

	if err := pipe.Set(videoIdS+"Index", tmpN, time.Hour*12).Err(); err != nil {
		return err
	}

	if _, err = pipe.Exec(); err != nil {
		return err
	}

	return nil

	// videoIdS := strconv.FormatInt(videoId, 10)
	// exist, err := CheckCommentsCache(videoId)
	// if err != nil {
	// 	return err
	// }
	// if !exist {
	// 	return errors.New("videoId is not exist in cache")
	// }
	// return RedisClient.Del(videoIdS).Err()
}

func UpdateCommentsCache(videoId int64, comment *comment.Comment) error {
	pipe := RedisClient.TxPipeline()

	// comment_cache change
	exist, err := CheckCommentsCache(videoId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("videoId is not exist in cache")
	}

	tmp, err := GetCommentsCache(videoId)
	if err != nil {
		return err
	}

	tmp = append(tmp, nil)
	copy(tmp[1:], tmp[0:])
	tmp[0] = comment

	videoIdS := strconv.FormatInt(videoId, 10)
	store, err := json.Marshal(tmp)
	if err != nil {
		return err
	}

	if err := pipe.Set(videoIdS, store, time.Hour*12).Err(); err != nil {
		return nil
	}

	// comment_index_cache change
	exist, err = CheckCommentIndexCache(videoId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("videoId is not exist in cache")
	}

	tmpN, err := GetCommentIndexCache(videoId)
	if err != nil {
		return err
	}
	tmpN += 1

	if err := pipe.Set(videoIdS+"Index", tmpN, time.Hour*12).Err(); err != nil {
		return err
	}

	if _, err = pipe.Exec(); err != nil {
		return err
	}

	return nil
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
