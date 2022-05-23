package main

import (
	douyin_user "github.com/douyin/kitex_gen/douyin_user/userservice"
	"log"
)

func main() {
	svr := douyin_user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
