package main

import (
	"example/cmd/api/handlers"
	"example/cmd/api/middleware"
	"example/cmd/api/rpc"
	"example/pkg/constants"
	"example/pkg/logger"
	"example/pkg/tracer"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Description
 * @Author 拥抱漏风
 * @Date 2022/5/20 17:34
 **/

func Init() {
	logger.Init()
	tracer.InitJaeger(constants.ApiServiceName)
	rpc.InitRPC()
}

func main() {
	Init()
	r := gin.New()
	r.POST("/douyin/user/register/", handlers.Register)
	r.POST("/douyin/user/login/", handlers.Login)
	r.GET("/douyin/feed/", handlers.VideoFeed)
	r.GET("/douyin/user/", handlers.UserInfo)
	r.Use(middleware.AuthMiddleWare())
	r.GET("/douyin/publish/list/", handlers.GetPublishVideo)
	r.GET("/douyin/favorite/list/", handlers.GetFavoriteVideo)
	r.POST("/douyin/publish/action/", handlers.PublishVideo)
	r.GET("/douyin/relation/follow/list/", handlers.GetFollowList)
	r.GET("/douyin/relation/follower/list/", handlers.GetFollowerList)
	r.POST("/douyin/relation/action/", handlers.RelationAction)
	r.POST("/douyin/favorite/action/", handlers.FavoriteAction)

	// 评论
	r.POST("/douyin/comment/action/", handlers.CommentAction)
	r.GET("/douyin/comment/list", handlers.GetComment)

	//r.GET("/douyin/auth", handlers.Identify)
	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}
