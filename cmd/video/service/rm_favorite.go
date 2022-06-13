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

type RmFavoriteService struct {
	ctx context.Context
}

func NewRmFavoriteService(ctx context.Context) *RmFavoriteService {
	return &RmFavoriteService{ctx: ctx}
}

func (s *RmFavoriteService) RmFavorite(req *video.RmFavoriteRequest) error {
	err := redispool.RemoveVideoId(req.UserId, req.VideoId)
	if err != nil {
		return err
	}
	return nil
}
