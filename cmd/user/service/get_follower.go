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
 * @Date 2022/5/23 17:03
 **/

type GetFollowerService struct {
	ctx context.Context
}

func NewGetFollowerService(ctx context.Context) *GetFollowerService {
	return &GetFollowerService{
		ctx: ctx,
	}
}

// GetFollower 获取粉丝列表
func (s *GetFollowerService) GetFollower(id int64) ([]*userdemo.User, error) {
	//fmt.Println("id:", id)
	userIds, err := db.GetFollowerIds(id)

	if err != nil {
		return make([]*userdemo.User, 0), err
	}
	//fmt.Println("userIds:", userIds)
	res := make([]*userdemo.User, 0)
	followerCount := 0
	userIdFollow, err := db.GetFollowerIds(id)
	if err == nil {
		followerCount = len(userIdFollow)
	}
	//fmt.Println("followerCount:", followerCount)

	// 更新用户粉丝数量(MySQL)
	_ = db.SetFollowerCountById(s.ctx, id, int64(followerCount))

	// 返回粉丝列表(redis)
	for _, val := range userIds {
		var dbUser *db.User
		userId, _ := strconv.ParseInt(string(val.([]byte)), 10, 64)
		dbUser, _ = db.GetUser(s.ctx, userId)
		demoUser := pack.User(id, dbUser)
		//demoUser.IsFollow = db.IsFollowerExist(id, userId)
		res = append(res, demoUser)
	}
	return res, nil
}
