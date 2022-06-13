package rpc

import (
	"context"
	"example/cmd/video/kitex_gen/video"
	"example/cmd/video/kitex_gen/video/videoservice"
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
 * @Date 2022/5/22 17:17
 **/

var videoClient videoservice.Client

func initVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
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
	videoClient = c
}

// GetVideoFeed 获取视频流
func GetVideoFeed(ctx context.Context, req *video.GetVideoFeedRequest) ([]*video.Video, error) {
	resp, err := videoClient.GetVideoFeed(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := resp.VideoList
	return res, nil
}

// PublishVideo 用户发布视频
func PublishVideo(ctx context.Context, req *video.CreateVideoRequest) (int64, error) {
	resp, err := videoClient.CreateVideo(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	res := resp.VideoId
	return res, nil
}

// AddFavorite 给视频点赞
func AddFavorite(ctx context.Context, req *video.AddFavoriteRequest) error {
	resp, err := videoClient.AddFavorite(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// RmFavorite 取消视频点赞
func RmFavorite(ctx context.Context, req *video.RmFavoriteRequest) error {
	resp, err := videoClient.RmFavorite(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// GetPublishVideo 获取用户发布的视频
func GetPublishVideo(ctx context.Context, req *video.GetPublishListRequest) ([]*video.Video, error) {
	resp, err := videoClient.GetPublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}

// GetFavoriteVideo 获取用户点赞视频
func GetFavoriteVideo(ctx context.Context, req *video.FavoriteVideoRequest) ([]*video.Video, error) {
	resp, err := videoClient.FavoriteVideo(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}
