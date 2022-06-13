package service

import (
	"context"
	"example/cmd/user/dal/db"
	"example/cmd/user/kitex_gen/userdemo"
	"example/pkg/errno"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 09:33
 **/

type CheckUserService struct {
	ctx context.Context
}

func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckUserService) CheckUser(req *userdemo.CheckUserRequest) (int64, error) {
	username := req.UserName
	users, err := db.QueryUser(s.ctx, username)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]
	if u.Password != req.Password {
		return 0, errno.LoginErr
	}
	return u.ID, nil
}
