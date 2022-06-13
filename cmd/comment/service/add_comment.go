package service

import (
	"context"
	"example/cmd/comment/dal/db"
	"example/cmd/comment/dal/redispool"
	"example/cmd/comment/kitex_gen/comment"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/30 10:46
 **/

type AddCommentService struct {
	ctx context.Context
}

func NewAddCommentService(ctx context.Context) *AddCommentService {
	return &AddCommentService{ctx: ctx}
}

func (s *AddCommentService) AddCommentService(req *comment.AddCommentRequest) error {
	singleComment := db.Comment{
		VideoId: req.VideoId,
		UserId:  req.UserId,
		Content: req.CommentText,
	}
	err := db.CreateComment(s.ctx, &singleComment)
	if err != nil {
		return err
	}
	err = redispool.IncrCommentCount(req.VideoId)
	if err != nil {
		return err
	}
	return nil
}
