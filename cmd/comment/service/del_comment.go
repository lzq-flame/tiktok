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
 * @Date 2022/5/30 14:31
 **/

type DelCommentService struct {
	ctx context.Context
}

func NewDelCommentService(ctx context.Context) *DelCommentService {
	return &DelCommentService{ctx: ctx}
}

func (s *DelCommentService) DelCommentService(req *comment.DelCommentRequest) error {
	singleComment := db.Comment{
		ID:      req.CommentId,
		VideoId: req.VideoId,
		UserId:  req.UserId,
		Content: "",
	}
	err := db.DeleteComment(s.ctx, &singleComment)
	if err != nil {
		return err
	}
	err = redispool.DecrCommentCount(req.VideoId)
	if err != nil {
		return err
	}
	return nil
}
