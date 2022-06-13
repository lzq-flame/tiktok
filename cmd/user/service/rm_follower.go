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
 * @Date 2022/5/28 14:47
 **/

type RmFollowerService struct {
	ctx context.Context
}

func NewRmFollowerService(ctx context.Context) *RmFollowerService {
	return &RmFollowerService{
		ctx: ctx,
	}
}

func (s *RmFollowerService) RmFollower(req *userdemo.RmFollowerRequest) error {
	redisFollowSetName := "user_follow_set_" + strconv.FormatInt(req.UserId, 10)
	redisFollowerSetName := "user_follower_set_" + strconv.FormatInt(req.FollowerId, 10)
	// 将toUserId在自己的关注列表中移除
	err := db.RemoveUserId(redisFollowSetName, req.FollowerId)
	if err != nil {
		return err
	}
	// 将自己在toUserId的粉丝列表中移除
	err = db.RemoveUserId(redisFollowerSetName, req.UserId)
	if err != nil {
		return err
	}
	return nil
}
