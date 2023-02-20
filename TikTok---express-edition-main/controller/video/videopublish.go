package video

import (
	"context"
	"douyin/service"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"strconv"
)

func Publish(c context.Context, ctx *app.RequestContext) {
	data, err := ctx.FormFile("data")
	userid := ctx.GetString("userid")
	Id, _ := strconv.ParseInt(userid, 10, 64)
	log.Printf("获取到用户id：%v", Id)
	title := ctx.PostForm("title")
	if err != nil {
		log.Println("获取数据错误", err)
	}
	Response := service.Publish(data, Id, title)
	ctx.JSON(int(Response.StatusCode), Response)
}
