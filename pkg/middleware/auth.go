package middleware

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/22 14:09
 **/

//func AuthMiddleWare() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		var resp handlers.UserRegisterResponse
//		//获取token
//		tokenString := ctx.PostForm("token")
//		if tokenString == "" {
//			tokenString = ctx.Query("token")
//		}
//		token, claims, err := jwt.ParseToken(tokenString)
//		if err != nil || !token.Valid {
//			resp.StatusCode = -1
//			resp.StatusMsg = "权限不足"
//			logger.L().Info("token 校验失败，权限不足")
//			ctx.JSON(http.StatusUnauthorized, &resp)
//			ctx.Abort()
//			return
//		}
//		userIds := make([]int64, 0)
//		userIds = append(userIds, claims.UserId)
//		userList, err := rpc.GetUsers(context.Background(), &userdemo.GetUserRequest{UserIds: userIds})
//		if err != nil {
//			resp.StatusCode = -1
//			resp.StatusMsg = "权限不足"
//			logger.L().Info("token 校验失败，权限不足")
//			ctx.JSON(http.StatusUnauthorized, &resp)
//			ctx.Abort()
//			return
//		}
//		if userList == nil || len(userList) == 0 {
//			resp.StatusCode = -1
//			resp.StatusMsg = "权限不足"
//			logger.L().Info("token 校验失败，权限不足")
//			ctx.JSON(http.StatusUnauthorized, &resp)
//			ctx.Abort()
//			return
//		}
//		//用户存在 将userId的值写入上下文
//		ctx.Set("userId", userList[0].Id)
//		logger.L().Info("token校验成功")
//		ctx.Next()
//	}
//}
