package main

import (
	"context"
	"example/cmd/user/kitex_gen/userdemo"
	"example/cmd/video/kitex_gen/video"
	"example/cmd/video/pack"
	"example/cmd/video/rpc"
	"example/cmd/video/service"
	"example/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *video.CreateVideoRequest) (resp *video.CreateVideoResp, err error) {
	resp = new(video.CreateVideoResp)
	id, err := service.NewCreateVideoService(ctx).CreateVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoId = id
	return resp, nil
}

// GetVideoFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoFeed(ctx context.Context, req *video.GetVideoFeedRequest) (resp *video.GetVideoFeedResp, err error) {
	resp = new(video.GetVideoFeedResp)
	videos, err := service.NewGetVideoFeedService(ctx).GetVideoFeed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}

// GetPublishList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPublishList(ctx context.Context, req *video.GetPublishListRequest) (resp *video.GetPublishListResp, err error) {
	resp = new(video.GetPublishListResp)
	videos, err := service.NewGetPublishVideo(ctx).GetPublishVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	user, err := rpc.GetUsers(ctx, &userdemo.GetUsersRequest{
		UserIds: []int64{req.UserId},
		MyId:    0,
	})
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	for _, v := range videos {
		v.User.UserName = user[req.UserId].UserName
		v.User.FollowerCount = user[req.UserId].FollowCount
		v.User.Id = req.UserId
		v.User.FollowCount = user[req.UserId].FollowCount
		v.User.IsFollow = user[req.UserId].IsFollow
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}

// AddFavorite implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) AddFavorite(ctx context.Context, req *video.AddFavoriteRequest) (resp *video.AddFavoriteResp, err error) {
	resp = new(video.AddFavoriteResp)
	err = service.NewAddFavoriteService(ctx).AddFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// RmFavorite implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) RmFavorite(ctx context.Context, req *video.RmFavoriteRequest) (resp *video.RmFavoriteResp, err error) {
	resp = new(video.RmFavoriteResp)
	err = service.NewRmFavoriteService(ctx).RmFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// FavoriteVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) FavoriteVideo(ctx context.Context, req *video.FavoriteVideoRequest) (resp *video.FavoriteVideoResp, err error) {
	resp = new(video.FavoriteVideoResp)
	videoList, err := service.NewGetFavoriteVideoService(ctx).GetFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.VideoList = videoList
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}
