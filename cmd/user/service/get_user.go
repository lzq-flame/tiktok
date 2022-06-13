package service

import (
	"context"
	"example/cmd/user/dal/db"
	"example/cmd/user/kitex_gen/userdemo"
	"example/cmd/user/pack"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 14:39
 **/

type GetUserService struct {
	ctx context.Context
}

func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

func (s *GetUserService) GetUsers(req *userdemo.GetUsersRequest) ([]*userdemo.User, error) {
	modelUsers, err := db.GetUsers(s.ctx, req.UserIds)
	if err != nil {
		return nil, err
	}
	return pack.Users(req.MyId, modelUsers), nil
}
