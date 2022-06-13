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
 * @Date 2022/5/28 14:28
 **/

// GetFollowerList 获取粉丝列表
func GetFollowerList(ctx *gin.Context) {
	var resp FollowerListResponse
	id, _ := ctx.Get("userId")
	userId, _ := id.(int64)
	followers, err := rpc.GetFollower(context.Background(), &userdemo.GetFollowerRequest{UserId: userId})
	if err != nil {
		logger.L().Info("调用 Get_Follow rpc失败, err : ", err)
		resp.StatusCode = -1
		resp.StatusMsg = "获取失败"
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}
	for _, v := range followers {
		resp.UserList = append(resp.UserList, PackUser(v))
	}
	resp.StatusCode = 0
	resp.StatusMsg = "获取成功"
	ctx.JSON(http.StatusOK, &resp)
}
