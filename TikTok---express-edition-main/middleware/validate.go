package middleware

import (
	"context"
	"douyin/pojo"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"github.com/go-ozzo/ozzo-validation/v3/is"
)

func LoginValidate(c context.Context, ctx *app.RequestContext) {
	username := ctx.Query("username")
	//validation校验用户名是否为空或者是否为一个邮箱
	err := validation.Validate(username, validation.Required.Error("username is required"), is.Email.Error("username must be a email address"))
	if err != nil {
		fmt.Println(err)
		ctx.Abort()
		ctx.JSON(200, pojo.UserLoginResponse{
			Response: pojo.Response{StatusCode: -1, StatusMsg: "用户名错误"},
		})
		return
	}
	ctx.Next(c)
	return
}
func RegisteryValidate(c context.Context, ctx *app.RequestContext) {
	username := ctx.Query("username")
	err1 := validation.Validate(username,
		validation.Required.Error("username is required"),
		is.Email.Error("userAccount必须是邮箱"))
	if err1 != nil {
		fmt.Println("userAccount必须是邮箱", err1)
		ctx.Abort()
		ctx.JSON(200, pojo.UserLoginResponse{
			Response: pojo.Response{StatusCode: -1, StatusMsg: "userAccount必须是邮箱"},
		})
		return
	}
	password := ctx.Query("password")
	err2 := validation.Validate(password, validation.Required.Error("password is required"), validation.RuneLength(6, 100).Error("字符长度应该在5~100之间"))
	if err2 != nil {
		fmt.Println("密码格式错误", err2)
		ctx.Abort()
		ctx.JSON(200, pojo.UserLoginResponse{
			Response: pojo.Response{StatusCode: -1, StatusMsg: "密码格式错误"},
		})
		return
	}
	ctx.Next(c)
	return
}
