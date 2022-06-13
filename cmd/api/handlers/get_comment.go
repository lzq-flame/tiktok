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
 * @Date 2022/5/30 15:33
 **/

func GetComment(ctx *gin.Context) {
	var req CommentListReq
	var resp CommentListResp
	//获取userId
	id, _ := ctx.Get("userId")
	userId, _ := id.(int64)
	if err := ctx.ShouldBind(&req); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "缺少必要的参数"
		logger.L().Info("请求参数缺失")
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}
	commentList, err := rpc.GetCommentList(context.Background(), &comment.GetCommentListRequest{
		VideoId: req.VideoId,
		UserId:  userId,
	})
	if err != nil {
		logger.L().Info("调用 GetCommentList rpc 失败 err: ", err)
		resp.StatusCode = -1
		resp.StatusMsg = "获取失败"
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}
	resp.CommentList = packCommentList(commentList)
	resp.StatusCode = 0
	resp.StatusMsg = "获取成功"
	ctx.JSON(http.StatusOK, &resp)
	return
}

func packCommentList(cs []*comment.Comment) []Comment {
	if cs == nil || len(cs) == 0 {
		return make([]Comment, 0)
	}
	ans := make([]Comment, len(cs))
	for k, v := range cs {
		ans[k].Id = v.CommentId
		ans[k].Content = v.Content
		ans[k].CreateDate = v.CreateDate
		if v.User == nil {
			ans[k].CommentUser = ProtoUser{}
			continue
		} else {
			ans[k].CommentUser.Name = v.User.UserName
			ans[k].CommentUser.FollowCount = v.User.FollowCount
			ans[k].CommentUser.FollowerCount = v.User.FollowerCount
			ans[k].CommentUser.IsFollow = v.User.IsFollow
			ans[k].CommentUser.ID = v.User.Id
		}
	}
	return ans
}
