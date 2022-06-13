package rpc

import (
	"context"
	"example/cmd/comment/kitex_gen/comment"
	"example/cmd/comment/kitex_gen/comment/commentservice"
	"example/pkg/constants"
	"example/pkg/errno"
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
 * @Date 2022/5/30 15:16
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

// AddComment 添加评论
func AddComment(ctx context.Context, req *comment.AddCommentRequest) error {
	resp, err := commentClient.AddComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// DelComment 删除评论
func DelComment(ctx context.Context, req *comment.DelCommentRequest) error {
	resp, err := commentClient.DelComment(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// GetCommentList 获取评论列表
func GetCommentList(ctx context.Context, req *comment.GetCommentListRequest) ([]*comment.Comment, error) {
	resp, err := commentClient.GetCommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.CommentList, nil
}
