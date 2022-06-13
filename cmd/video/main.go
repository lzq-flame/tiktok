package main

import (
	"example/cmd/video/dal"
	video "example/cmd/video/kitex_gen/video/videoservice"
	"example/cmd/video/rpc"
	"example/pkg/bound"
	"example/pkg/constants"
	"example/pkg/logger"
	"example/pkg/middleware"
	tracer2 "example/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"net"
)

func Init() {
	logger.Init()
	rpc.InitRPC()
	tracer2.InitJaeger(constants.VideoServiceName)
	dal.Init()
}

func main() {
	r, err := etcd.NewEtcdRegistry([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	if err != nil {
		panic(err)
	}
	Init()
	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: constants.VideoServiceName}), // server name
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
