package handlers

import (
	"example/cmd/api/rpc"
	"example/cmd/user/kitex_gen/userdemo"
	"example/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/29 09:22
 **/

func UserInfo(ctx *gin.Context) {
	var req UserInfoReq
	var resp UserInfoResp
	if err := ctx.ShouldBind(&req); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "缺少参数"
		logger.L().Info("缺少id或token")
		ctx.JSON(http.StatusPreconditionFailed, &resp)
		return
	}

	user, err := rpc.GetUser(ctx, &userdemo.GetUsersRequest{
		UserIds: []int64{req.UserId},
		MyId:    0,
	})
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "获取用户信息失败"
		logger.L().Info("调用 GetUser rpc 失败 err: ", err)
		ctx.JSON(http.StatusPreconditionFailed, &resp)
		return
	}
	resp.User.FollowCount = user.FollowCount
	resp.User.IsFollow = user.IsFollow
	resp.User.ID = user.Id
	resp.User.FollowerCount = user.FollowerCount
	resp.User.Name = user.UserName
	resp.StatusCode = 0
	resp.StatusMsg = "获取成功"
	ctx.JSON(http.StatusOK, &resp)
	return
}
