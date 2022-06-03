package handlers

import (
	"context"
	"douyin/cmd/api/rpc"
	"douyin/kitex_gen/comment"
	"douyin/kitex_gen/user"
	"douyin/middleware"
	"douyin/pkg/errno"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	rpc.Response
	CommentList []*rpc.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	rpc.Response
	Comment rpc.Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	_, claims, err := middleware.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, &rpc.Response{StatusCode: errno.CommentServiceErrCode, StatusMsg: err.Error()})
		return
	}
	userId := claims.UserID

	_, err = rpc.QueryUser(context.Background(), &user.DouyinUserRequest{
		UserId: userId,
		Token:  token,
	})
	if err != nil {
		c.JSON(http.StatusOK, rpc.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	if actionType == "1" {
		text := c.Query("comment_text")
		videoIdS := c.Query("video_id")
		videoId, err := strconv.ParseInt(videoIdS, 10, 64)

		fmt.Println(videoIdS)

		if err != nil {
			c.JSON(http.StatusOK, &rpc.Response{StatusCode: errno.ServiceErrCode, StatusMsg: err.Error()})
			return
		}

		response, comment := rpc.CreateComment(context.Background(), &comment.CreateCommentRequest{
			UserId:  userId,
			VideoId: videoId,
			Content: text,
		})
		if response.StatusCode != 0 {
			c.JSON(http.StatusOK, response)
			return
		}

		c.JSON(http.StatusOK, CommentActionResponse{Response: *response, Comment: *comment})
		return
	} else if actionType == "2" {
		videoIdS := c.Query("video_id")
		videoId, err := strconv.ParseInt(videoIdS, 10, 64)

		if err != nil {
			c.JSON(http.StatusOK, &rpc.Response{StatusCode: errno.ServiceErrCode, StatusMsg: err.Error()})
			return
		}

		commentIdS := c.Query("comment_id")
		commentId, err := strconv.ParseInt(commentIdS, 10, 64)

		if err != nil {
			c.JSON(http.StatusOK, &rpc.Response{StatusCode: errno.ServiceErrCode, StatusMsg: err.Error()})
			return
		}

		response := rpc.DeleteComment(context.Background(), &comment.DeleteCommentRequest{
			UserId:    userId,
			VideoId:   videoId,
			CommentId: commentId,
		})

		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, rpc.Response{StatusCode: errno.ActionTypeErr.ErrCode, StatusMsg: errno.ActionTypeErr.ErrMsg})
}

func CommentList(c *gin.Context) {
	videoIdS := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdS, 10, 64)
	token := c.Query("token")

	if err != nil {
		c.JSON(http.StatusOK, &rpc.Response{StatusCode: errno.ServiceErrCode, StatusMsg: err.Error()})
		return
	}

	response, comments := rpc.QueryComments(context.Background(), &comment.QueryCommentsRequest{VideoId: videoId, Token: &token})

	if response.StatusCode != 0 {
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    *response,
		CommentList: comments,
	})
}
