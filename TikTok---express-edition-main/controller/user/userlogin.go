package user

import (
	"context"
	"douyin/pojo"
	"douyin/service"

	"github.com/cloudwego/hertz/pkg/app"
)

var LoginResponse pojo.UserLoginResponse

func UserLogin(c context.Context, ctx *app.RequestContext) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	LoginResponse := service.Iflogin(c, username, password)
	ctx.JSON(int(LoginResponse.Response.StatusCode), LoginResponse)
}
