package redisdb

import (
	"douyin/kitex_gen/comment"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type Comments_cache struct {
	CommentNumber int64              `json:"comment_number"`
	Comments      []*comment.Comment `json:"comments,omitempty"`
}

func AddCommentsCache(videoId int64, comments []*comment.Comment) error {
	videoIdS := strconv.FormatInt(videoId, 10)

	comments_cache := Comments_cache{
		CommentNumber: int64(len(comments)),
		Comments:      comments,
	}

	store, err := json.Marshal(comments_cache)
	if err != nil {
		return err
	}
	return RedisClient.Set(videoIdS, store, time.Hour*12).Err()
}

func AddCommentNumberCache(videoId, CommentNumber int64) error {
	videoIdS := strconv.FormatInt(videoId, 10)

	comments_cache := Comments_cache{
		CommentNumber: CommentNumber,
		Comments:      nil,
	}

	store, err := json.Marshal(comments_cache)
	if err != nil {
		return err
	}
	return RedisClient.Set(videoIdS, store, time.Hour*12).Err()
}

func DeleteCommentsCache(videoId, commentId int64) error {
	pipe := RedisClient.TxPipeline()

	// comment_cache change
	exist, tmp, err := CheckGetCommentsCache(videoId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("videoId is not exist in cache")
	}

	if err != nil {
		return err
	}

	index := 0
	for k, value := range tmp.Comments {
		if value.CommentId == commentId {
			index = k
			break
		}
	}

	tmp.Comments = append(tmp.Comments[:index], tmp.Comments[index+1:]...)
	tmp.CommentNumber -= 1

	videoIdS := strconv.FormatInt(videoId, 10)
	store, err := json.Marshal(tmp)
	if err != nil {
		return err
	}

	if err := pipe.Set(videoIdS, store, time.Hour*12).Err(); err != nil {
		return nil
	}

	if _, err = pipe.Exec(); err != nil {
		return err
	}

	return nil
}

func UpdateCommentsCache(videoId int64, comment *comment.Comment) error {
	pipe := RedisClient.TxPipeline()

	// comment_cache change
	exist, tmp, err := CheckGetCommentsCache(videoId)
	if err != nil {
		return err
	}
	if !exist {
		return errors.New("videoId is not exist in cache")
	}

	if err != nil {
		return err
	}

	tmp.Comments = append(tmp.Comments, nil)
	copy(tmp.Comments[1:], tmp.Comments[0:])
	tmp.Comments[0] = comment
	tmp.CommentNumber += 1

	videoIdS := strconv.FormatInt(videoId, 10)
	store, err := json.Marshal(tmp)
	if err != nil {
		return err
	}

	if err := pipe.Set(videoIdS, store, time.Hour*12).Err(); err != nil {
		return nil
	}

	if _, err = pipe.Exec(); err != nil {
		return err
	}

	return nil
}

func CheckGetCommentsCache(videoId int64) (bool, *Comments_cache, error) {
	videoIdS := strconv.FormatInt(videoId, 10)
	exist, err := RedisClient.Exists(videoIdS).Result()

	if err != nil {
		return false, nil, err
	}

	if exist <= 0 {
		return false, nil, nil
	}

	resS, err := RedisClient.Get(videoIdS).Result()
	if err != nil {
		return false, nil, err
	}
	var res *Comments_cache
	if err = json.Unmarshal([]byte(resS), &res); err != nil {
		return false, nil, err
	}

	if res.CommentNumber != int64(len(res.Comments)) {
		return false, nil, nil
	}

	return true, res, err
}

func CheckCommentNumberCache(videoId int64) (bool, error) {
	videoIdS := strconv.FormatInt(videoId, 10)
	exist, err := RedisClient.Exists(videoIdS).Result()

	if err != nil {
		return false, err
	}

	if exist <= 0 {
		return false, nil
	}

	return true, err
}

func GetCommentIndexCache(videoId int64) (int64, error) {
	videoIdS := strconv.FormatInt(videoId, 10)
	resS, err := RedisClient.Get(videoIdS).Result()
	if err != nil {
		return -1, err
	}
	var res *Comments_cache
	if err = json.Unmarshal([]byte(resS), &res); err != nil {
		return -1, err
	}
	return res.CommentNumber, nil
}
