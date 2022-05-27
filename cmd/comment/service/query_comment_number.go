package service

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/cmd/comment/dal/redisdb"
	"douyin/kitex_gen/comment"
	"log"
	"strconv"

	"golang.org/x/sync/singleflight"
)

var gQueryCommentNumber singleflight.Group

type QueryCommentNumberService struct {
	ctx context.Context
}

func NewQueryCommentNumberService(ctx context.Context) *QueryCommentNumberService {
	return &QueryCommentNumberService{ctx: ctx}
}

func (s *QueryCommentNumberService) QueryCommentNumber(req *comment.QueryCommentNumberRequest) (int64, error) {
	status, err := redisdb.CheckCommentIndexCache(req.VideoId)

	if err != nil {
		log.Fatal("redis err: ", err)
	}

	if status {
		res, err := redisdb.GetCommentIndexCache(req.VideoId)
		if err == nil {
			return res, err
		}
		log.Fatal("redis err: ", err)
	}

	key := strconv.FormatInt(req.VideoId, 10)

	v, err, _ := gQueryCommentNumber.Do(key, func() (interface{}, error) {
		res, err := mysqldb.QueryCommentsNumber(s.ctx, req.VideoId)
		if err != nil {
			return 0, err
		}
		redisdb.AddCommentIndexCache(req.VideoId, res)
		return res, nil
	})

	if err != nil {
		return 0, err
	}

	return v.(int64), nil
}
