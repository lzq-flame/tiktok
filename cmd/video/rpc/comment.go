package rpc

import (
	"context"
	"example/cmd/comment/kitex_gen/comment"
	"example/cmd/comment/kitex_gen/comment/commentservice"
	"example/pkg/constants"
	"example/pkg/middleware"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/30 15:52
 **/

var commentClient commentservice.Client

func initCommentRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := commentservice.NewClient(
		constants.CommentServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	commentClient = c
}

func GetCommentCount(ctx context.Context, req *comment.GetCommentCountRequest) int64 {
	resp, err := commentClient.GetCommentCount(ctx, req)
	if err != nil {
		return 0
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0
	}
	return resp.CommentCount
}
