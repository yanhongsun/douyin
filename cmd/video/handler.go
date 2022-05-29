package main

import (
	"context"
	"douyin/cmd/video/assist"
	"douyin/cmd/video/dal/db"
	"douyin/kitex_gen/video"
	"douyin/pkg/constants"
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
	db.PublishVideo(ctx, video)
	sucess := "sucess"
	resp.BaseResp.StatusCode = 0
	resp.BaseResp.StatusMsg = &sucess
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
	//TODO
	userId := req.UserId
	videos, _ := db.GetPublishList(ctx, userId)
	//为返回值赋值
	//TODO根据Token得到userId
	var meId int64
	if req.Token != "" {
		meId, err = strconv.ParseInt(req.Token, 10, 64)
		if err != nil {
			fmt.Println("handler.go->strconv.ParseInt error")
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
			fmt.Println("handler.go->strconv.ParseInt error")
			fmt.Println(err)
			return
		}
	} else {

		userId = constants.EmptyUserId
	}

	videos, err := db.GetFeed(ctx, *(req.LatestTime), constants.VideoLimitNum)
	//为返回值赋值

	r := assist.GetFeedResp(ctx, videos, userId)

	resp = &r
	return
}
