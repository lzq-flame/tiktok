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
 * @Date 2022/5/20 17:35
 **/

func Register(ctx *gin.Context) {
	var req UserRegisterRequest
	var resp UserRegisterResponse
	if err := ctx.ShouldBind(&req); err != nil {
		resp.StatusCode = errno.ConvertErr(err).ErrCode
		resp.StatusMsg = "缺少用户名或密码"
		logger.L().Info("缺少用户名或密码", resp)
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}
	if len(req.Username) == 0 || len(req.Username) > 32 || len(req.Password) == 0 || len(req.Password) > 32 {
		resp.StatusCode = errno.ParamErrCode
		resp.StatusMsg = "用户名或密码长度不正确"
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}

	id, err := rpc.CreateUser(context.Background(), &userdemo.CreateUserRequest{
		UserName: req.Username,
		Password: req.Password,
	})
	if err != nil {
		resp.StatusCode = errno.ConvertErr(err).ErrCode
		resp.StatusMsg = errno.ConvertErr(err).ErrMsg
		ctx.JSON(http.StatusBadRequest, &resp)
		return
	}
	//生成token
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
