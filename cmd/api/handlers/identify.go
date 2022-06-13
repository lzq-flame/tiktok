package handlers

import (
	"context"
	"example/cmd/api/rpc"
	"example/cmd/user/kitex_gen/userdemo"
	"example/pkg/jwt"
	"example/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 14:33
 **/

func Identify(ctx *gin.Context) {
	var resp UserRegisterResponse
	//获取token
	tokenString := ctx.PostForm("token")
	if tokenString == "" {
		tokenString = ctx.Query("token")
	}
	token, claims, err := jwt.ParseToken(tokenString)
	fmt.Println(tokenString)
	if err != nil || !token.Valid {
		resp.StatusCode = -1
		resp.StatusMsg = "权限不足"
		logger.L().Info("1.token 校验失败，权限不足")
		ctx.JSON(http.StatusUnauthorized, &resp)
		ctx.Abort()
		return
	}
	userIds := make([]int64, 0)
	userIds = append(userIds, claims.UserId)
	//logger.L().Info("userIds:", userIds)
	userList, err := rpc.GetUsers(context.Background(), &userdemo.GetUsersRequest{UserIds: userIds})
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "权限不足"
		logger.L().Info("2.token 校验失败，权限不足")
		ctx.JSON(http.StatusUnauthorized, &resp)
		ctx.Abort()
		return
	}
	if userList == nil || len(userList) == 0 {
		resp.StatusCode = -1
		resp.StatusMsg = "权限不足"
		logger.L().Info("3.token 校验失败，权限不足")
		ctx.JSON(http.StatusUnauthorized, &resp)
		ctx.Abort()
		return
	}
	//用户存在 将userId的值写入上下文
	//ctx.Set("userId", userList[0].Id)
	ctx.Header("x-user", strconv.Itoa(int(userList[0].Id)))
	logger.L().Info("token校验成功")
	ctx.Next()
}
