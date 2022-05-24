package main

import (
	"douyin/cmd/comment/dal"
	comment "douyin/kitex_gen/comment/commentservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})

	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")

	if err != nil {
		panic(err)
	}

	Init()

	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "comment"}),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 1000}),
		server.WithMuxTransport(),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
