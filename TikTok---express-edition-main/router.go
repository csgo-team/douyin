package main

import (
	"context"
	"douyin/controller/comment"
	"douyin/controller/user"
	"douyin/controller/video"
	"douyin/middleware"
	"douyin/pojo"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func initrouter(h *server.Hertz) {
	hg := h.Group("/douyin")
	{
		//注册
		hg.POST("/user/register/", middleware.RegisteryValidate, user.UserRegister)

		//登录
		hg.POST("/user/login/", middleware.LoginValidate, user.UserLogin)

		//信息展示
		hg.GET("/user/", middleware.Auth, user.UserInfo)
		//投稿
		hg.POST("/publish/action/", middleware.AuthPublish, video.Publish)
		////视频流接口
		hg.GET("/feed/", video.Feed)
		//用户视频列表
		hg.GET("/publish/list/", middleware.Auth, video.PublishList)

		//评论操作
		hg.POST("/comment/action/", middleware.Auth, comment.CommentAction)
		//评论列表
		hg.GET("/comment/list/", middleware.Auth, comment.CommentList)
		// 点赞接口
		hg.POST("favorite/action/", middleware.Auth, video.IsLike)
	}
	println("hertz初始化完成")
}

// 测试feed流接口所用函数以及demodata
func handler_1(c context.Context, ctx *app.RequestContext) {
	fmt.Println("视频流")
	DemoUser := pojo.User{
		Id:            7,
		Name:          "demouser",
		FollowCount:   2,
		FollowerCount: 2,
		IsFollow:      false,
	}
	//http://localhost:8888/video/testvideo.m3u8
	vids := []pojo.Video{
		{
			Id:            699962082483127216,
			Author:        DemoUser,
			PlayUrl:       "https://csgovideo.oss-cn-guangzhou.aliyuncs.com/video/2023February9a252909e-9d4a-417e-9ca2-3b0cf5247467.mp4",
			CoverUrl:      "https://csgovideo.oss-cn-guangzhou.aliyuncs.com/cover/test.jpg",
			FavoriteCount: 1,
			CommentCount:  1,
			IsFavorite:    true,
			Title:         "demo",
		},
		{
			Id:            699962082483127217,
			Author:        DemoUser,
			PlayUrl:       "https://csgovideo.oss-cn-guangzhou.aliyuncs.com/video/2023February9cfd7117d-89ef-4579-8cb4-cbb24fc6802a.mp4",
			CoverUrl:      "https://csgovideo.oss-cn-guangzhou.aliyuncs.com/cover/test.jpg",
			FavoriteCount: 1,
			CommentCount:  1,
			IsFavorite:    true,
			Title:         "demo",
		},
	}
	ctx.JSON(200, pojo.FeedResponse{Response: pojo.Response{
		StatusCode: 0,
		StatusMsg:  "欢迎观看",
	}, NextTime: time.Now().Unix(),
		VideoList: vids,
	})
}
