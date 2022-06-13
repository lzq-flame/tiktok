package handlers

import (
	"context"
	"example/cmd/api/rpc"
	"example/cmd/user/kitex_gen/userdemo"
	"example/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/28 13:53
 **/

// GetFollowList 获取关注列表
func GetFollowList(ctx *gin.Context) {
	var resp FollowListResponse

	id, _ := ctx.Get("userId")
	userId, _ := id.(int64)
	follows, err := rpc.GetFollow(context.Background(), &userdemo.GetFollowRequest{UserId: userId})
	if err != nil {
		logger.L().Info("调用 Get_Follow rpc失败, err : ", err)
		resp.StatusCode = -1
		resp.StatusMsg = "获取失败"
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}
	for _, v := range follows {
		resp.UserList = append(resp.UserList, PackUser(v))
	}
	resp.StatusCode = 0
	resp.StatusMsg = "获取成功"
	ctx.JSON(http.StatusOK, &resp)
}

func PackUser(user *userdemo.User) ProtoUser {
	if user == nil {
		return ProtoUser{}
	}
	return ProtoUser{
		ID:            user.Id,
		Name:          user.UserName,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}
