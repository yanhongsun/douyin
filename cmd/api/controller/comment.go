package controller

import (
	"context"
	"douyin/cmd/api/common"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/comment"
	"douyin/pkg/errno"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	common.Response
	CommentList []*common.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	common.Response
	Comment common.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			videoIdS := c.Query("video_id")
			videoId, err := strconv.ParseInt(videoIdS, 10, 64)

			if err != nil {
				c.JSON(http.StatusOK, &common.Response{StatusCode: errno.ServiceErrCode, StatusMsg: err.Error()})
			}

			response, comment := rpc.CreateComment(context.Background(), &comment.CreateCommentRequest{
				UserId:  user.Id,
				VideoId: videoId,
				Content: text,
			})
			if response.StatusCode != 0 {
				c.JSON(http.StatusOK, response)
			}
			c.JSON(http.StatusOK, CommentActionResponse{Response: *response, Comment: *comment})
			return
		} else if actionType == "2" {
			videoIdS := c.Query("video_id")
			videoId, err := strconv.ParseInt(videoIdS, 10, 64)

			if err != nil {
				c.JSON(http.StatusOK, &common.Response{StatusCode: errno.ServiceErrCode, StatusMsg: err.Error()})
			}

			commentIdS := c.Query("comment_id")
			commentId, err := strconv.ParseInt(commentIdS, 10, 64)

			if err != nil {
				c.JSON(http.StatusOK, &common.Response{StatusCode: errno.ServiceErrCode, StatusMsg: err.Error()})
			}

			response := rpc.DeleteComment(context.Background(), &comment.DeleteCommentRequest{
				UserId:    user.Id,
				VideoId:   videoId,
				CommentId: commentId,
			})
			c.JSON(http.StatusOK, response)
			return
		}
		c.JSON(http.StatusOK, common.Response{StatusCode: errno.ActionTypeErr.ErrCode, StatusMsg: errno.ActionTypeErr.ErrMsg})
	} else {
		c.JSON(http.StatusOK, common.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoIdS := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdS, 10, 64)

	if err != nil {
		c.JSON(http.StatusOK, &common.Response{StatusCode: errno.ServiceErrCode, StatusMsg: err.Error()})
	}

	response, comments := rpc.QueryComments(context.Background(), &comment.QueryCommentsRequest{VideoId: videoId})

	if response.StatusCode != 0 {
		c.JSON(http.StatusOK, response)
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    *response,
		CommentList: comments,
	})
}
