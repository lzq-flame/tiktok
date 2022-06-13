package service

import (
	"context"
	"example/cmd/video/dal/redispool"
	"example/cmd/video/kitex_gen/video"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/28 17:28
 **/

type AddFavoriteService struct {
	ctx context.Context
}

func NewAddFavoriteService(ctx context.Context) *AddFavoriteService {
	return &AddFavoriteService{ctx: ctx}
}

func (s *AddFavoriteService) AddFavorite(req *video.AddFavoriteRequest) error {
	err := redispool.AddVideoId(req.UserId, req.VideoId)
	if err != nil {
		return err
	}
	return nil
}
