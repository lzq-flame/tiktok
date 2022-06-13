package handlers

import (
	"example/cmd/api/rpc"
	"example/cmd/video/kitex_gen/video"
	"example/pkg/errno"
	"example/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 15:05
 **/

func PublishVideo(ctx *gin.Context) {
	var resp PublishActionResp
	id, _ := ctx.Get("userId")
	userId, _ := id.(int64)
	file, err := ctx.FormFile("data")
	title := ctx.PostForm("title")
	if err != nil {
		logger.L().Warn("获取data失败")
		resp.StatusCode = -1
		resp.StatusMsg = "获取视频数据失败"
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}

	videoId, err := rpc.PublishVideo(ctx, &video.CreateVideoRequest{
		UserId: userId,
		Title:  title,
	})
	if err != nil {
		logger.L().Warn("调用rpc失败")
		resp.StatusCode = int32(errno.ConvertErr(err).ErrCode)
		resp.StatusMsg = errno.ConvertErr(err).ErrMsg
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}

	//保存视频
	vid := strconv.FormatInt(videoId, 10)
	filename := vid + ".mp4"
	err = ctx.SaveUploadedFile(file, "../../data/video/"+filename)
	if err != nil {
		logger.L().Warn("保存视频失败")
		resp.StatusCode = int32(errno.ConvertErr(err).ErrCode)
		resp.StatusMsg = errno.ConvertErr(err).ErrMsg
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}
	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = "Success"
	resp.VideoId = videoId
	ctx.JSON(http.StatusOK, &resp)
	return
}
