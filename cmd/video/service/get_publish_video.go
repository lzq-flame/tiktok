package service

import (
	"context"
	"example/cmd/video/dal/db"
	"example/cmd/video/dal/redispool"
	"example/cmd/video/kitex_gen/video"
	"example/cmd/video/pack"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/28 16:05
 **/

type GetPublishVideoService struct {
	ctx context.Context
}

func NewGetPublishVideo(ctx context.Context) *GetPublishVideoService {
	return &GetPublishVideoService{ctx: ctx}
}

func (s *GetPublishVideoService) GetPublishVideo(req *video.GetPublishListRequest) ([]*video.Video, error) {
	videoIds, err := redispool.GetVideoIdsByUserId(req.UserId)
	if err != nil {
		return nil, err
	}
	videos, err := db.GetVideoListByVideoIds(s.ctx, videoIds)
	if err != nil {
		return nil, err
	}
	return pack.Videos(req.UserId, videos), nil
}
