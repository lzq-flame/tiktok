package handlers

import (
	"context"
	"example/cmd/api/rpc"
	"example/cmd/video/kitex_gen/video"
	"example/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/29 09:12
 **/

func FavoriteAction(ctx *gin.Context) {
	var req FavoriteActionRequest
	var resp FavoriteActionResponse
	ctx.ShouldBind(&req)
	id, _ := ctx.Get("userId")
	userId, _ := id.(int64)
	if req.VideoId == 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "缺少请求参数"
		logger.L().Info("视频id不存在")
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}
	if req.ActionType == 1 {
		err := rpc.AddFavorite(context.Background(), &video.AddFavoriteRequest{
			UserId:  userId,
			VideoId: req.VideoId,
		})
		if err != nil {
			resp.StatusCode = -1
			resp.StatusMsg = "点赞失败"
			logger.L().Info("调用AddFavorite rpc失败 err: ", err)
			ctx.JSON(http.StatusBadRequest, &resp)
			return
		}
		resp.StatusMsg = "点赞成功"
	} else {
		err := rpc.RmFavorite(ctx, &video.RmFavoriteRequest{
			UserId:  userId,
			VideoId: req.VideoId,
		})
		if err != nil {
			resp.StatusCode = -1
			resp.StatusMsg = "取消点赞失败"
			logger.L().Info("调用RmFavorite rpc失败 err: ", err)
			ctx.JSON(http.StatusBadRequest, &resp)
			return
		}
		resp.StatusMsg = "取消点赞"
	}
	resp.StatusCode = 0
	ctx.JSON(http.StatusOK, &resp)
	return
}
