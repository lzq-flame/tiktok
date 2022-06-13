package rpc

import (
	"context"
	"example/cmd/user/kitex_gen/userdemo"
	"example/cmd/user/kitex_gen/userdemo/userservice"
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
 * @Date 2022/5/20 17:48
 **/

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
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
	userClient = c
}

// CreateUser create user info
func CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (int64, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Id, nil
}

// CheckUser check user info
func CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}

// GetUsers 获取用户信息列表
func GetUsers(ctx context.Context, req *userdemo.GetUsersRequest) ([]*userdemo.User, error) {
	resp, err := userClient.GetUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := resp.Users
	return res, nil
}

// GetUser 获取用户信息列表
func GetUser(ctx context.Context, req *userdemo.GetUsersRequest) (*userdemo.User, error) {
	resp, err := userClient.GetUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := resp.Users[0]
	return res, nil
}

// GetFollow 获取关注列表
func GetFollow(ctx context.Context, req *userdemo.GetFollowRequest) ([]*userdemo.User, error) {
	resp, err := userClient.GetFollow(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Follows, nil
}

// GetFollower 获取粉丝列表
func GetFollower(ctx context.Context, req *userdemo.GetFollowerRequest) ([]*userdemo.User, error) {
	resp, err := userClient.GetFollower(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Followers, nil
}

// AddFollow 添加关注
func AddFollow(ctx context.Context, req *userdemo.AddFollowerRequest) error {
	resp, err := userClient.AddFollower(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// RmFollow 取消关注
func RmFollow(ctx context.Context, req *userdemo.RmFollowerRequest) error {
	resp, err := userClient.RmFollower(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}
