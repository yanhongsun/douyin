package pack

import (
	"context"
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/cmd/comment/rpc"
	"douyin/kitex_gen/comment"
	"time"
)

func ChangeComment(source *mysqldb.Comment, user *rpc.UserInfo) *comment.Comment {
	return &comment.Comment{
		CommentId: source.CommentID,
		User: &comment.User{
			Id:            user.ID,
			Name:          user.Name,
			FollowCount:   &user.FollowCount,
			FollowerCount: &user.FollowerCount,
			IsFollow:      user.IsFollow,
		},
		Content:    source.Content,
		CreateDate: time.Now().Format("01-02"),
	}
}

func ChangeComments(ctx context.Context, source []*mysqldb.Comment, token *string) ([]*comment.Comment, error) {
	size := len(source)
	res := make([]*comment.Comment, size)

	for i := 0; i < size; i++ {
		user, err := rpc.GetUserInfo(ctx, source[i].UserID, *token)
		if err != nil {
			return nil, err
		}
		res[i] = ChangeComment(source[i], user)
	}

	return res, nil
}
