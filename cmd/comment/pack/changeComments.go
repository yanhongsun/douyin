package pack

import (
	"douyin/cmd/comment/dal/mysqldb"
	"douyin/kitex_gen/comment"
	"time"
)

func ChangeComment(source *mysqldb.Comment) *comment.Comment {
	return &comment.Comment{
		CommentId:  source.CommentID,
		UserId:     source.UserID,
		Content:    source.Content,
		CreateDate: time.Now().Format("01-02"),
	}
}

func ChangeComments(source []*mysqldb.Comment) []*comment.Comment {
	size := len(source)
	res := make([]*comment.Comment, size)

	for i := 0; i < size; i++ {
		res[i] = ChangeComment(source[i])
	}

	return res
}
