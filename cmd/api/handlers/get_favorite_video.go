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
 * @Date 2022/5/30 09:23
 **/

func GetFavoriteVideo(ctx *gin.Context) {
	var resp FavoriteListResponse
	id, _ := ctx.Get("userId")
	userId, _ := id.(int64)
	videoList, err := rpc.GetFavoriteVideo(context.Background(), &video.FavoriteVideoRequest{
		UserId: userId,
	})
	if err != nil {
		logger.L().Info("调用 GetFavoriteVideo rpc 失败 err: ", err)
		resp.StatusCode = -1
		resp.StatusMsg = "获取失败"
		ctx.JSON(http.StatusInternalServerError, &resp)
		return
	}
	for _, v := range videoList {
		resp.VideoList = append(resp.VideoList, PackVideo(v))
	}
	resp.StatusCode = 0
	resp.StatusMsg = "Success"
	ctx.JSON(http.StatusOK, &resp)
	return
}
