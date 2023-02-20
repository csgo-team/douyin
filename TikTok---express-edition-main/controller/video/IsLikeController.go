package video

import (
	"context"
	"douyin/pojo"
	"douyin/service"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// IsLike 点赞接口 /**
func IsLike(c context.Context, ctx *app.RequestContext) {
	// 从上下文中获取 ActionType 状态，
	isLikeStr := ctx.Query("ActionType")
	idStr := ctx.Query("user_id")
	videoIdStr := ctx.Query("VideoId")
	if len(isLikeStr) == 0 || len(idStr) == 0 || len(videoIdStr) == 0 {
		ctx.JSON(-1, pojo.Response{
			StatusCode: -1,
			StatusMsg:  "参数错误",
		})
	}

	isLikeNum, err := strconv.Atoi(isLikeStr)
	//idNum, err := strconv.Atoi(idStr)
	//videoIdNum, err := strconv.Atoi(videoIdStr)
	if err != nil {
		ctx.JSON(-1, pojo.Response{
			StatusCode: -1,
			StatusMsg:  "参数转化错误",
		})
	}

	if isLikeNum != 1 || isLikeNum != 0 {
		ctx.JSON(-1, pojo.Response{
			StatusCode: -1,
			StatusMsg:  "点赞状态错误",
		})
	}

	response := service.DoLike(c, videoIdStr, idStr, int64(isLikeNum))
	ctx.JSON(int(response.StatusCode), response)

}
