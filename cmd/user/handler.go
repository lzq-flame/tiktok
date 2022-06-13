package main

import (
	"context"
	"example/cmd/user/kitex_gen/userdemo"
	"example/cmd/user/pack"
	"example/cmd/user/service"
	"example/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (resp *userdemo.CreateUserResp, err error) {
	resp = new(userdemo.CreateUserResp)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	id, err := service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Id = id
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (resp *userdemo.CheckUserResp, err error) {
	resp = new(userdemo.CheckUserResp)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil

}

// AddFollower implements the UserServiceImpl interface.
func (s *UserServiceImpl) AddFollower(ctx context.Context, req *userdemo.AddFollowerRequest) (resp *userdemo.AddFollowerResp, err error) {
	resp = new(userdemo.AddFollowerResp)
	err = service.NewAddFollowerService(ctx).AddFollower(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// RmFollower implements the UserServiceImpl interface.
func (s *UserServiceImpl) RmFollower(ctx context.Context, req *userdemo.RmFollowerRequest) (resp *userdemo.RmFollowerResp, err error) {
	resp = new(userdemo.RmFollowerResp)
	err = service.NewRmFollowerService(ctx).RmFollower(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// GetFollow implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollow(ctx context.Context, req *userdemo.GetFollowRequest) (resp *userdemo.GetFollowResp, err error) {
	resp = new(userdemo.GetFollowResp)
	demoUsers := make([]*userdemo.User, 0)
	demoUsers, err = service.NewGetFollowService(ctx).GetFollow(req.UserId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Follows = demoUsers
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// GetFollower implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetFollower(ctx context.Context, req *userdemo.GetFollowerRequest) (resp *userdemo.GetFollowerResp, err error) {
	resp = new(userdemo.GetFollowerResp)
	demoUsers := make([]*userdemo.User, 0)
	//fmt.Println("req:", req)
	demoUsers, err = service.NewGetFollowerService(ctx).GetFollower(req.UserId)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Followers = demoUsers
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return
}

// GetUsers implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUsers(ctx context.Context, req *userdemo.GetUsersRequest) (resp *userdemo.GetUsersResp, err error) {
	resp = new(userdemo.GetUsersResp)
	if len(req.UserIds) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}
	var users []*userdemo.User
	users, err = service.NewGetUserService(ctx).GetUsers(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.Users = users
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
