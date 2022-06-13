package handlers

import (
	"context"
	"example/cmd/api/rpc"
	"example/cmd/video/kitex_gen/video"
	"example/pkg/errno"
	"example/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/23 16:24
 **/

func VideoFeed(ctx *gin.Context) {
	var timeStamp int64
	var resp VideoFeedResp
	timeStamp = time.Now().Unix()

	// 获取视频列表
	videos, err := rpc.GetVideoFeed(context.Background(), &video.GetVideoFeedRequest{
		TimeStamp: timeStamp,
	})
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "获取失败"
		logger.L().Warn("调用 GetVideoFeed rpc 失败 err:%v ", err)
		ctx.JSON(http.StatusInternalServerError, &resp)
		return
	}
	for _, v := range videos {
		resp.VideoList = append(resp.VideoList, PackVideo(v))
	}
	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = "Success"
	ctx.JSON(http.StatusOK, &resp)
	return
}

func PackVideo(v *video.Video) Video {
	if v == nil {
		return Video{}
	}
	hv := Video{
		ID:            v.Id,
		Title:         v.Title,
		Author:        ProtoUser{},
		PlayUrl:       v.PlayUrl,
		CoverUrl:      v.CoverUrl,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    v.IsFavorite,
	}
	if v.User == nil {
		return hv
	}
	hv.Author.IsFollow = v.User.IsFollow
	hv.Author.ID = v.User.Id
	hv.Author.Name = v.User.UserName
	hv.Author.FollowCount = v.User.FollowCount
	hv.Author.FollowerCount = v.User.FollowerCount
	return hv
}
