package main

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"github.com/yanhongsun/douyin/cmd/thumb/dal"
	like "github.com/yanhongsun/douyin/kitex_gen/like/thumbservice"
	"github.com/yanhongsun/douyin/pkg/bound"
	"github.com/yanhongsun/douyin/pkg/constants"
	"github.com/yanhongsun/douyin/pkg/middleware"
	"github.com/yanhongsun/douyin/pkg/tracer"
	"net"
)

func Init() {
	tracer.InitJaeger(constants.ThumbServiceName)
	dal.Init()
}
func main() {
	Init()
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}
	svr := like.NewServer(new(ThumbServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.ThumbServiceName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                              // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()),                // BoundHandler
		server.WithRegistry(r),                                             // registry
	)

	err = svr.Run()

	if err != nil {
		klog.Fatal(err)
	}
}
