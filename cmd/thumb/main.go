package main

import (
	like "github.com/yanhongsun/douyin/kitex_gen/like/thumbservice"
	"log"
)

func main() {
	svr := like.NewServer(new(ThumbServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
