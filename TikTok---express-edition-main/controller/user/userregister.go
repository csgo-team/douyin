package user

import (
	"context"
	"douyin/pojo"
	"douyin/service"
	"github.com/cloudwego/hertz/pkg/app"
)

var RegisterResponse pojo.UserLoginResponse

func UserRegister(c context.Context, ctx *app.RequestContext) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	RegisterResponse = service.IfRegister(username, password)
	ctx.JSON(200, RegisterResponse)
}
