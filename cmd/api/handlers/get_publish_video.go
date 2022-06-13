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
 * @Date 2022/5/28 16:55
 **/

func GetPublishVideo(ctx *gin.Context) {
	var resp PublishListResp
	id, _ := ctx.Get("userId")
	userId, _ := id.(int64)
	videoList, err := rpc.GetPublishVideo(context.Background(), &video.GetPublishListRequest{
		UserId: userId,
	})
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "获取失败"
		logger.L().Info("调用 GetPublishVideo rpc 失败 err: ", err)
		ctx.JSON(http.StatusBadRequest, &resp)
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
