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
 * @Date 2022/5/20 16:49
 **/

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser 创建用户
func (s *CreateUserService) CreateUser(req *userdemo.CreateUserRequest) (int64, error) {
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return 0, err
	}
	if len(users) != 0 {
		return 0, errno.UserAlreadyExistErr
	}

	user := db.User{
		Username: req.UserName,
		Password: req.Password,
	}

	err = db.CreateUser(s.ctx, &user)

	return user.ID, err

}
