package pack

import (
	"douyin/cmd/comment/dal/db"
	"douyin/kitex_gen/comment"
)

func ChangeComment(source *db.Comment) *comment.Comment {
	return &comment.Comment{
		CommentId:  source.CommentID,
		UserId:     source.UserID,
		Content:    source.Content,
		CreateDate: source.CreatedAt.Format("01-02"),
	}
}

func ChangeComments(source []*db.Comment) []*comment.Comment {
	size := len(source)
	res := make([]*comment.Comment, size)

	for i := 0; i < size; i++ {
		res[i] = ChangeComment(source[i])
	}

	return res
}
