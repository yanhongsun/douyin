package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/douyin/cmd/user/config"
	"github.com/douyin/cmd/user/dal"
	"github.com/douyin/cmd/user/global"
	middleware2 "github.com/douyin/cmd/user/middleware"
	user "github.com/douyin/kitex_gen/user/userservice"
	"github.com/douyin/middleware"
	"github.com/douyin/pkg/bound"
	"github.com/douyin/pkg/tracer"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"log"
	"net"
)

func Init() {
	tracer.InitJaeger(global.ServerSetting.UserServName)

	dal.Init()

	err := setupSetting()
	if err != nil {
		log.Println(err)
	}

	setupJWT()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{fmt.Sprintf("%s:%s",
		global.ServerSetting.EtcdAddress,
		global.ServerSetting.EtcdPort)})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s",
		global.ServerSetting.UserServAddr,
		global.ServerSetting.UserServPort))
	if err != nil {
		panic(err)
	}

	Init()
	// TODO: add server configuration
	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: global.ServerSetting.UserServName}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                                     // middleware
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
		// log.Println(err.Error())
		klog.Fatal(err)
	}
}

// init functions
// setupSetting initialize global settings from yaml
func setupSetting() error {
	setting, err := config.NewSetting()
	if err != nil {
		return err
	}

	// load database config
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	// load JWT config
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	return nil
}

// setupJWT
func setupJWT() {
	global.Jwt = middleware2.NewJWT()
}
