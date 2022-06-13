package handlers

import (
	"context"
	"example/cmd/api/rpc"
	"example/cmd/comment/kitex_gen/comment"
	"example/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/30 15:26
 **/

func CommentAction(ctx *gin.Context) {
	var req CommentRequest
	var resp CommentResponse
	if err := ctx.ShouldBind(&req); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "缺少必要的参数"
		logger.L().Info("请求参数缺失")
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}
	//获取userId
	id, _ := ctx.Get("userId")
	userId, _ := id.(int64)
	if req.ActionType == 1 {
		if len(req.CommentText) == 0 {
			resp.StatusCode = -1
			resp.StatusMsg = "评论内容为空"
			ctx.JSON(http.StatusBadRequest, &resp)
			return
		}
		err := rpc.AddComment(context.Background(), &comment.AddCommentRequest{
			UserId:      userId,
			VideoId:     req.VideoId,
			CommentText: req.CommentText,
		})
		if err != nil {
			resp.StatusCode = -1
			resp.StatusMsg = "评论失败"
			logger.L().Info("调用addComment rpc 失败 err: ", err)
			ctx.JSON(http.StatusBadRequest, &resp)
			return
		}
		resp.StatusMsg = "评论成功"
	} else {
		err := rpc.DelComment(context.Background(), &comment.DelCommentRequest{
			CommentId: req.CommentId,
			UserId:    userId,
			VideoId:   req.VideoId,
		})
		if err != nil {
			resp.StatusCode = -1
			resp.StatusMsg = "删除评论失败"
			logger.L().Info("DelComment rpc 失败 err: ", err)
			ctx.JSON(http.StatusBadRequest, &resp)
			return
		}
		resp.StatusMsg = "删除评论成功"
	}
	resp.StatusCode = 0
	ctx.JSON(http.StatusOK, &resp)
	return
}
