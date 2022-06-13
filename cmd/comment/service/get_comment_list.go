package service

import (
	"context"
	"example/cmd/comment/dal/db"
	"example/cmd/comment/kitex_gen/comment"
	"example/cmd/comment/pack"
	"example/cmd/comment/rpc"
	"example/cmd/user/kitex_gen/userdemo"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/30 14:54
 **/

type GetCommentListService struct {
	ctx context.Context
}

func NewGetCommentListService(ctx context.Context) *GetCommentListService {
	return &GetCommentListService{ctx: ctx}
}

func (s *GetCommentListService) GetCommentList(req *comment.GetCommentListRequest) ([]*comment.Comment, error) {
	comments, err := db.GetCommentByVideoId(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0)
	for _, c := range comments {
		ids = append(ids, c.UserId)
	}
	userMap, err := rpc.GetUsers(context.Background(), &userdemo.GetUsersRequest{
		UserIds: ids,
		MyId:    req.UserId,
	})
	if err != nil {
		return nil, err
	}
	users := make([]*userdemo.User, 0)
	for i := 0; i < len(ids); i++ {
		users = append(users, userMap[ids[i]])
	}
	return pack.Comments(users, comments), nil
}
