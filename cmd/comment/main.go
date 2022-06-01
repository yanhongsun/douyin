package main

import (
	"douyin/cmd/comment/dal"
	"douyin/cmd/comment/pack/configdata"
	"douyin/cmd/comment/repository"
	"douyin/cmd/comment/rpc"
	"douyin/cmd/comment/service"
	comment "douyin/kitex_gen/comment/commentservice"

	tracer2 "douyin/pkg/tracer"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	err := configdata.SetupSetting()
	if err != nil {
		panic("config is wrong")
	}

	tracer2.InitJaegers(configdata.CommentServerConfig.CommentServName)
	rpc.InitRPC()
	dal.Init()
	repository.Init()
	service.InitSnowflakeNode()
}

func main() {
	Init()

	r, err := etcd.NewEtcdRegistry([]string{configdata.CommentServerConfig.EtcdHost})

	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", configdata.CommentServerConfig.CommentServHost)

	if err != nil {
		panic(err)
	}

	svr := comment.NewServer(new(CommentServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: configdata.CommentServerConfig.CommentServName}),
		server.WithServiceAddr(addr),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 1000}),
		server.WithMuxTransport(),
		server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
