package pack

import (
	"douyin/cmd/comment/dal/mysqldb"
)

func ReverseComments(comments []*mysqldb.Comment) []*mysqldb.Comment {
	size := len(comments)
	res := make([]*mysqldb.Comment, size)
	count := 0
	for i := size - 1; i >= 0; i-- {
		res[count] = comments[i]
		count++
	}
	return res
}
