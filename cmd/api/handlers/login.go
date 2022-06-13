package handlers

import (
	"context"
	"example/cmd/api/rpc"
	"example/cmd/user/kitex_gen/userdemo"
	"example/pkg/errno"
	"example/pkg/jwt"
	"example/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 10:27
 **/

// Login 用户登录
func Login(ctx *gin.Context) {
	var req UserLoginRequest
	var resp UserLoginResponse
	if err := ctx.ShouldBind(&req); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "缺少用户名或密码"
		logger.L().Info("userinfo 缺少用户名或密码 %v", req)
		ctx.JSON(http.StatusOK, &resp)
		return
	}

	id, err := rpc.CheckUser(context.Background(), &userdemo.CheckUserRequest{
		UserName: req.Username,
		Password: req.Password,
	})
	if err != nil {
		resp.StatusCode = errno.ConvertErr(err).ErrCode
		resp.StatusMsg = errno.ConvertErr(err).ErrMsg
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}
	token, err := jwt.ReleaseToken(id)
	if err != nil {
		logger.L().Warn("生成token失败")
		resp.StatusCode = -1
		resp.StatusMsg = "生成token失败"
		ctx.JSON(http.StatusInternalServerError, &resp)
		return
	}
	resp.UserId = id
	resp.Token = token
	resp.StatusCode = errno.SuccessCode
	resp.StatusMsg = "Success"
	ctx.JSON(http.StatusOK, &resp)
	return
}
