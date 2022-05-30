package main

import (
	"context"
	"douyin/cmd/video/assist"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/video"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"strconv"

	"fmt"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishVideoRequest) (resp *video.PublishVideoResponse, err error) {
	//TODO从token解析用户ID
	//TODO异常处理
	var userId int64
	userId = 1 //req.Token
	//验证userId是否存在
	//如果不存在返回err
	// resp.BaseResp.StatusCode = errno.UserNotExistErr.ErrCode
	// resp.BaseResp.StatusMsg = &errno.UserNotExistErr.ErrMsg
	//err=nil
	if len(req.Data) == 0 {
		resp.BaseResp.StatusCode = errno.VideoErr.ErrCode
		resp.BaseResp.StatusMsg = &errno.VideoErr.ErrMsg
		err = nil
		return
	}
	//将视频数据命名，保存路径
	name := GetName(userId)
	playUrl := fmt.Sprintf("%s/%s.mp4", constants.VideoSavePath, name)
	assist.SaveVideo(playUrl, req.Data)
	//从视频数据中提取封面，保存路径
	coverUrl := fmt.Sprintf("%s/%s.png", constants.VideoCoverSavePath, name)
	assist.GetCover(coverUrl, playUrl)
	//为db.Video类型变量赋值
	//为db.Video类型变量写进数据库
	video := assist.InitDBVideo(userId, playUrl, coverUrl, req.Title)
	if err = db.PublishVideo(ctx, video); err == nil {
		resp.BaseResp.StatusCode = errno.Success.ErrCode
		resp.BaseResp.StatusMsg = &errno.Success.ErrMsg
	}

	return
}

//userId+当前时间
func GetName(userId int64) string {
	return fmt.Sprintf("%s%s", string(userId), time.Now().Format("2006-01-02 15:04:05"))
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *video.GetPublishListRequest) (resp *video.GetPublishListResponse, err error) {
	// TODO: Your code here..
	//TODO异常处理.
	//TODO验证token
	userId := req.UserId
	videos, err := db.GetPublishList(ctx, userId)
	if err != nil {
		return nil, err
	}
	//为返回值赋值
	//TODO根据Token得到userId
	var meId int64
	if req.Token != "" {
		meId, err = strconv.ParseInt(req.Token, 10, 64)
		if err != nil {
			fmt.Println("handlers.go->strconv.ParseInt error")
			fmt.Println(err)
			return
		}
	} else {

		meId = constants.EmptyUserId
	}
	r := assist.GetPublishListResp(ctx, videos, meId)
	resp = &r
	return resp, nil
}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *video.GetFeedRequest) (resp *video.GetFeedResponse, err error) {
	// TODO: Your code here...
	//异常处理
	//token处理
	//token处理
	var userId int64
	if *req.Token != "" {
		userId, err = strconv.ParseInt(*req.Token, 10, 64)
		if err != nil {
			fmt.Println("handlers.go->strconv.ParseInt error")
			fmt.Println(err)
			return
		}
	} else {

		userId = constants.EmptyUserId
	}

	videos, err := db.GetFeed(ctx, *(req.LatestTime), constants.VideoLimitNum)
	//为返回值赋值
	if err != nil {
		return nil, err
	}
	r := assist.GetFeedResp(ctx, videos, userId)

	resp = &r
	return
}

// VerifyVideoId implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) VerifyVideoId(ctx context.Context, req *video.VerifyVideoIdRequest) (resp *video.VerifyVideoIdResponse, err error) {

	//TODO: 验证token
	var tOrf bool
	if req.VideoId < 0 {
		tOrf = false
		err = nil
	} else {
		tOrf, err = db.VerifyVideoId(ctx, req.VideoId)
	}
	if err != nil {
		return nil, err
	}
	r := assist.VerifyVideoIdResp(tOrf)
	resp = &r
	return
}
