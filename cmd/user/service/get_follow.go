package service

import (
	"context"
	"example/cmd/user/dal/db"
	"example/cmd/user/kitex_gen/userdemo"
	"example/cmd/user/pack"
	"strconv"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/28 13:48
 **/

type GetFollowService struct {
	ctx context.Context
}

func NewGetFollowService(ctx context.Context) *GetFollowService {
	return &GetFollowService{
		ctx: ctx,
	}
}

// GetFollow 获取关注列表
func (s *GetFollowService) GetFollow(id int64) ([]*userdemo.User, error) {
	userIds, err := db.GetFollowIds(id)

	if err != nil {
		return make([]*userdemo.User, 0), err
	}
	res := make([]*userdemo.User, 0)
	followCount := 0
	userIdFollow, err := db.GetFollowIds(id)
	if err == nil {
		followCount = len(userIdFollow)
	}

	// 更新用户关注数量(MySQL)
	_ = db.SetFollowCountById(s.ctx, id, int64(followCount))

	// 返回关注列表(redis)
	for _, val := range userIds {
		var dbUser *db.User
		userId, _ := strconv.ParseInt(string(val.([]byte)), 10, 64)
		dbUser, _ = db.GetUser(s.ctx, userId)
		demoUser := pack.User(0, dbUser)
		demoUser.IsFollow = true
		res = append(res, demoUser)
	}
	return res, nil
}
