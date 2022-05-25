package main

import (
	user "github.com/douyin/kitex_gen/user/userservice"
	"log"
)

func main() {
	// TODO: add server configuration
	svr := user.NewServer(new(UserServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
