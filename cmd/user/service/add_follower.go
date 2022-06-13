package service

import (
	"context"
	"example/cmd/user/dal/db"
	"example/cmd/user/kitex_gen/userdemo"
	"strconv"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/23 16:47
 **/

type AddFollowerService struct {
	ctx context.Context
}

func NewAddFollowerService(ctx context.Context) *AddFollowerService {
	return &AddFollowerService{
		ctx: ctx,
	}
}

func (s *AddFollowerService) AddFollower(req *userdemo.AddFollowerRequest) error {
	redisFollowSetName := "user_follow_set_" + strconv.FormatInt(req.UserId, 10)
	redisFollowerSetName := "user_follower_set_" + strconv.FormatInt(req.FollowerId, 10)
	// 添加toUserId到自己的关注列表
	err := db.AddUserId(redisFollowSetName, req.FollowerId)
	if err != nil {
		return err
	}
	// 在toUserId的粉丝列表添加自己
	err = db.AddUserId(redisFollowerSetName, req.UserId)
	if err != nil {
		return err
	}
	return nil
}
