package service

import (
	"context"
	"example/cmd/comment/dal/redispool"
	"example/cmd/comment/kitex_gen/comment"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/30 14:47
 **/

type CommentCountService struct {
	ctx context.Context
}

func NewCommentCountService(ctx context.Context) *CommentCountService {
	return &CommentCountService{ctx: ctx}
}

func (s *CommentCountService) CommentCountService(req *comment.GetCommentCountRequest) int64 {
	return redispool.GetCommentCount(req.VideoId)
}
