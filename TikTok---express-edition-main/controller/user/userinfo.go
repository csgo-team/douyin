package user

import (
	"context"
	"douyin/pojo"
	"douyin/service"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

// UserInfo 通过登陆的用户在redis中的数据获取用户
func UserInfo(c context.Context, ctx *app.RequestContext) {
	idstr2 := ctx.Query("user_id")
	if len(idstr2) == 0 {
		ctx.JSON(-1, pojo.Response{
			StatusCode: -1,
			StatusMsg:  "没有用户信息",
		})
	}

	// 获取用户id
	id, _ := strconv.Atoi(idstr2)

	response := service.GetUserById(c, int64(id))
	ctx.JSON(int(response.StatusCode), response)
}
