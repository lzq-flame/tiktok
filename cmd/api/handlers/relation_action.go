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
 * @Date 2022/5/28 15:07
 **/

func RelationAction(ctx *gin.Context) {
	var req RelationActionRequest
	var resp RelationActionResponse
	var err error
	err = ctx.ShouldBind(&req)
	userId, _ := ctx.Get("userId")
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "缺少请求参数"
		logger.L().Info("RelationActionHandle 缺少请求参数 ")
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}

	if req.ActionType == 1 {
		err = rpc.AddFollow(context.Background(), &userdemo.AddFollowerRequest{
			UserId:     userId.(int64),
			FollowerId: req.ToUserId,
		})
		if err != nil {
			logger.L().Info("调用AddFollow rpc 失败 err: ", err)
			resp.StatusCode = -1
			resp.StatusMsg = "关注失败"
			ctx.JSON(http.StatusInternalServerError, &resp)
			return
		}
		resp.StatusMsg = "关注成功"
	} else {
		err = rpc.RmFollow(context.Background(), &userdemo.RmFollowerRequest{
			UserId:     userId.(int64),
			FollowerId: req.ToUserId,
		})
		if err != nil {
			logger.L().Info("调用RmFollow rpc 失败 err: ", err)
			resp.StatusCode = -1
			resp.StatusMsg = "失败"
			ctx.JSON(http.StatusInternalServerError, &resp)
			return
		}
		resp.StatusMsg = "取消关注"
	}
	resp.StatusCode = 0
	ctx.JSON(200, &resp)
	return
}
