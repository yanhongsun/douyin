package main

import (
	"douyin/cmd/comment/dal"
	"douyin/cmd/comment/pack/configdata"
	"douyin/cmd/comment/pack/zapcomment"
	"douyin/cmd/comment/repository"
	"douyin/cmd/comment/rpc"
	"douyin/cmd/comment/service"
	comment "douyin/kitex_gen/comment/commentservice"

	tracer2 "douyin/pkg/tracer"
	"net"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	zapcomment.Init()
	err := configdata.SetupSetting()
	if err != nil {
		zapcomment.Logger.Panic("configuration initialization failed")
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
		zapcomment.Logger.Panic("etcd initialization err: " + err.Error())
	}

	addr, err := net.ResolveTCPAddr("tcp", configdata.CommentServerConfig.CommentServHost)

	if err != nil {
		zapcomment.Logger.Panic("tcp address initialization err: " + err.Error())
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
		zapcomment.Logger.Panic("comment service initialization err: " + err.Error())
	}
}
