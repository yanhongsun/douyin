package main

import (
	"douyin/kitex_gen/video"
	"fmt"
	"time"
)

// import (
// 	"context"
// 	"douyin/cmd/video/assist"
// 	"douyin/cmd/video/dal/db"
// 	"douyin/kitex_gen/video"
// 	"douyin/pkg/constants"
// 	"fmt"
// 	"time"
// )

// VideoServiceImpl implements the last service interface defined in the IDL.

// PublishVideo implements the VideoServiceImpl interface.
func PublishVideo(req *video.PublishVideoRequest) (resp *video.PublishVideoResponse, err error) {
	//TODO从token解析用户ID
	//TODO异常处理
	//userId := req.Token
	//将视频数据命名，保存路径
	// name := assist.GetName(userId)
	// playUrl := fmt.Sprintf("%s/%s.mp4", constants.VideoSavePath, name)
	// assist.SaveVideo(playUrl, req.Data)
	// //从视频数据中提取封面，保存路径
	// coverUrl := fmt.Sprintf("%s/%s.png", constants.VideoCoverSavePath, name)
	// assist.GetCover(coverUrl, playUrl)
	// //为db.Video类型变量赋值
	// //为db.Video类型变量写进数据库
	// video := assist.InitDBVideo(userId, playUrl, coverUrl, req.Title)
	// db.PublishVideo(ctx, video)
	//
	// resp.BaseResp.StatusCode = 0
	// resp.BaseResp.StatusMsg = "sucess"
	resp = nil
	return resp, nil

}

//userId+当前时间
func GetName(userId int64) {
	fmt.Sprintf("%s%s", string(userId), time.Now().Format("2006-01-02 15:04:05"))
}

// GetPublishList implements the VideoServiceImpl interface.
// func GetPublishList(ctx context.Context, req *video.GetPublishListRequest) (resp *video.GetPublishListResponse, err error) {
// 	// TODO: Your code here..
// 	//TODO异常处理.
// 	userId := req.UserId
// 	videos, err := db.GetPublishList(ctx, userId)
// 	//为返回值赋值
// 	resp = assist.GetPublishListResp(videos, userId)
// 	return resp, nil
// }

// // GetFeed implements the VideoServiceImpl interface.
// func GetFeed(ctx context.Context, req *video.GetFeedRequest) (resp *video.GetFeedResponse, err error) {
// 	// TODO: Your code here...
// 	//异常处理
// 	//token处理
// 	userId := req.Token
// 	videos, err := db.GetFeed(ctx, req.LastTime, constants.VideoLimitNum)
// 	//为返回值赋值
// 	if userId == nil {
// 		resp = assist.GetFeedResp(videos, -1)
// 	} else {
// 		resp = assist.GetFeedResp(videos, userId)
// 	}

// 	return
// }

func main() {
	fmt.Println("sjiaIJ")
}
