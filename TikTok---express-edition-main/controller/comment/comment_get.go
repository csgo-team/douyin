package comment

import (
	"context"
	"douyin/pojo"
	"douyin/service"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"strconv"
)

var CommentGet pojo.CommontListResponse

func CommentList(c context.Context, ctx *app.RequestContext) {
	video_id := ctx.Query("video_id")
	id, err := strconv.Atoi(video_id)
	if err != nil {
		log.Println("字符转换数字出错", err)
		return
	}
	CommentGet := service.GetCommentbyID(int64(id))
	if CommentGet == nil {
		ctx.JSON(404, pojo.UserResponse{
			Response: pojo.Response{StatusMsg: "用户加载失败", StatusCode: -1},
		})
		return
	}
	ctx.JSON(200, pojo.UserResponse{
		Response: pojo.Response{StatusCode: 0, StatusMsg: "评论加载成功"},
	})
}
