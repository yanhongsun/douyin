package redisdb

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type Comments_cache struct {
	CommentNumber int64              `json:"comment_number"`
	Comments      []*mysqldb.Comment `json:"comments,omitempty"`
}

func AddCommentsCache(ctx context.Context, videoId int64, comments []*mysqldb.Comment) error {
	videoIdS := strconv.FormatInt(videoId, 10)

	comments_cache := Comments_cache{
		CommentNumber: int64(len(comments)),
		Comments:      comments,
	}

	store, err := json.Marshal(comments_cache)
	if err != nil {
		return err
	}
	return RedisClient.Set(ctx, videoIdS, store, time.Hour*12).Err()
}

func AddCommentNumberCache(ctx context.Context, videoId, CommentNumber int64) error {
	videoIdS := strconv.FormatInt(videoId, 10)

	comments_cache := Comments_cache{
		CommentNumber: CommentNumber,
		Comments:      nil,
	}

	store, err := json.Marshal(comments_cache)
	if err != nil {
		return err
	}
	return RedisClient.Set(ctx, videoIdS, store, time.Hour*12).Err()
}

func DeleteCommentsCache(ctx context.Context, videoId, commentId int64) error {
	exist, tmp, err := CheckCommentsCache(ctx, videoId)
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
		if value.CommentID == commentId {
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

	if err := RedisClient.Set(ctx, videoIdS, store, time.Hour*12).Err(); err != nil {
		return nil
	}

	return nil
}

func UpdateCommentsCache(ctx context.Context, videoId int64, comment *mysqldb.Comment) error {
	exist, tmp, err := CheckCommentsCache(ctx, videoId)
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

	if err := RedisClient.Set(ctx, videoIdS, store, time.Hour*12).Err(); err != nil {
		return nil
	}

	return nil
}

func CheckCommentsCache(ctx context.Context, videoId int64) (bool, *Comments_cache, error) {
	videoIdS := strconv.FormatInt(videoId, 10)
	exist, err := RedisClient.Exists(ctx, videoIdS).Result()

	if err != nil {
		return false, nil, err
	}

	if exist <= 0 {
		return false, nil, nil
	}

	resS, err := RedisClient.Get(ctx, videoIdS).Result()
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

func CheckCommentNumberCache(ctx context.Context, videoId int64) (bool, error) {
	videoIdS := strconv.FormatInt(videoId, 10)
	exist, err := RedisClient.Exists(ctx, videoIdS).Result()

	if err != nil {
		return false, err
	}

	if exist <= 0 {
		return false, nil
	}

	return true, err
}

func GetCommentIndexCache(ctx context.Context, videoId int64) (int64, error) {
	videoIdS := strconv.FormatInt(videoId, 10)
	resS, err := RedisClient.Get(ctx, videoIdS).Result()
	if err != nil {
		return -1, err
	}
	var res *Comments_cache
	if err = json.Unmarshal([]byte(resS), &res); err != nil {
		return -1, err
	}
	return res.CommentNumber, nil
}
