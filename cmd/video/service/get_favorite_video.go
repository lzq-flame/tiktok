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
 * @Date 2022/5/30 08:45
 **/

type GetFavoriteVideoService struct {
	ctx context.Context
}

func NewGetFavoriteVideoService(ctx context.Context) *GetFavoriteVideoService {
	return &GetFavoriteVideoService{ctx: ctx}
}

func (s *GetFavoriteVideoService) GetFavorite(req *video.FavoriteVideoRequest) ([]*video.Video, error) {
	ids, err := redispool.GetFavoriteVideo(req.UserId)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	videos, err := db.GetVideoListByVideoIds(s.ctx, ids)
	if err != nil {
		return nil, err
	}
	return pack.Videos(req.UserId, videos), nil
}
