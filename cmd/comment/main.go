package main

import (
	"douyin/cmd/comment/dal"
	comment "douyin/kitex_gen/comment/commentservice"
	"log"
)

func Init() {
	dal.Init()
}

func main() {
	Init()

	svr := comment.NewServer(new(CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
