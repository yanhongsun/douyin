package main

import (
	"context"
	"douyin/cmd/video/assist"
	"douyin/cmd/video/dal/db"
	"douyin/cmd/video/dal/minicache"
	"douyin/kitex_gen/video"
	"douyin/middleware"
	"douyin/pkg/constants"
	"douyin/pkg/errno"
	"fmt"
	"strconv"
	"time"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishVideoRequest) (resp *video.PublishVideoResponse, err error) {
	//TODO从token解析用户ID
	//TODO异常处理
	fmt.Println("Publish Video")
	var userId int64
	//req.Token
	//验证userId是否存在
	//如果不存在返回err
	// resp.BaseResp.StatusCode = errno.UserNotExistErr.ErrCode
	// resp.BaseResp.StatusMsg = &errno.UserNotExistErr.ErrMsg
	//err=nil
	resp = &video.PublishVideoResponse{}
	if req.Token != "" {
		_, claims, err := middleware.ParseToken(req.Token)
		if err != nil {
			return nil, err
		}
		userId = claims.UserID
	} else {
		userId = constants.EmptyUserId
	}
	if len(req.Data) == 0 {
		bresp := assist.GetBaseResp(errno.VideoErr.ErrCode, errno.VideoErr.ErrMsg)
		resp.BaseResp = &bresp
		err = nil
		return
	}

	//将视频数据命名，保存路径
	name := GetName(userId)

	playUrl := fmt.Sprintf("%s/%s/%s.mp4", constants.VideoResourceIpPort, constants.VideoUrlPath, name)
	videoSavePath := fmt.Sprintf("%s/%s.mp4", constants.VideoSavePath, name)
	assist.SaveVideo(videoSavePath, req.Data)
	//从视频数据中提取封面，保存路径
	coverUrl := fmt.Sprintf("%s/%s/%s.png", constants.VideoResourceIpPort, constants.VideoCoverUrlPath, name)
	coverSavePath := fmt.Sprintf("%s/%s.png", constants.VideoCoverSavePath, name)
	assist.GetCover(coverSavePath, videoSavePath)
	//为db.Video类型变量赋值
	//为db.Video类型变量写进数据库
	video := assist.InitDBVideo(userId, playUrl, coverUrl, req.Title)

	if err = minicache.PublishVideo(ctx, video); err == nil {
		//fmt.Println(errno.Success.ErrCode)
		bresp := assist.GetBaseResp(errno.Success.ErrCode, errno.Success.ErrMsg)
		resp.BaseResp = &bresp
	}

	return
}

//userId+当前时间
func GetName(userId int64) string {
	return fmt.Sprintf("%s%s", strconv.FormatInt(userId, 10), strconv.FormatInt(time.Now().UnixNano()/1e6, 10))
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *video.GetPublishListRequest) (resp *video.GetPublishListResponse, err error) {
	// TODO: Your code here..
	//TODO异常处理.
	//TODO验证token
	userId := req.UserId
	videos, err := minicache.GetPublishList(ctx, userId)
	if err != nil {
		return nil, err
	}
	//为返回值赋值
	//TODO根据Token得到userId
	var meId int64
	if req.Token != "" {
		_, claims, err := middleware.ParseToken(req.Token)
		if err != nil {
			return nil, err
		}
		meId = claims.UserID
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
		_, claims, err := middleware.ParseToken(*req.Token)
		if err != nil {
			return nil, err
		}
		userId = claims.UserID
	} else {
		userId = constants.EmptyUserId
	}

	videos, err := minicache.GetFeed(ctx, *(req.LatestTime), constants.VideoLimitNum)
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
