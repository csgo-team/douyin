package video

import (
	"context"
	"douyin/service"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func PublishList(c context.Context, ctx *app.RequestContext) {
	userId := ctx.Query("user_id")
	userid, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println("id类型转换出错", err)
		return
	}
	res := service.PublishList(userid)
	ctx.JSON(200, res)
}
