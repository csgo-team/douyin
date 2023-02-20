package middleware

import (
	"context"
	"douyin/Init/Redis"
	"douyin/config"
	"douyin/pojo"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strconv"
	"time"
)

var key = []byte("whjlwjwpflllzjq")

type userclaim struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *userclaim, error) {
	Claims := &userclaim{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		//key是设定的密钥
		return key, nil
	})
	return token, Claims, err
}

// Auth 健全逻辑，从redis中获取
func Auth(c context.Context, ctx *app.RequestContext) {
	redis := Redis.InitRedis()
	userid := ctx.Query("user_id")
	tokenKey := config.TOKEN_KEY + userid
	result, err2 := redis.Get(c, tokenKey).Result()
	if err2 != nil {
		errors.New("获取错误")
	}

	if len(result) == 0 {
		fmt.Println("没有token")
		ctx.Abort()
		ctx.JSON(200, pojo.Response{
			StatusCode: -1,
			StatusMsg:  "Unauthorized",
		})
		return
	}
	//解析token得到结构体并检查是否有效
	token, userclaim, err := ParseToken(result)
	if err != nil {
		fmt.Println("token解析出错")
		ctx.Abort()
		ctx.JSON(200, pojo.Response{
			StatusCode: -1,
			StatusMsg:  "Unauthorized",
		})
		return
	}
	if !token.Valid {
		//token无效,继续请求
		fmt.Println("token失效")
		ctx.Abort()
		ctx.JSON(http.StatusUnauthorized, pojo.Response{
			StatusCode: -1,
			StatusMsg:  "会话过期，请重新登录",
		})
		return
	} else {
		//token有效向上下文中加入
		redis.Set(c, tokenKey, result, time.Minute*10)
		ctx.Set("user_id", userclaim.UserId)
		ctx.Next(c)
	}
	return
}

// SetToken 颁发Token
func SetToken(user pojo.User) (tokenstr string) {
	//设定过期时间，这里是一个小时
	expirtime := time.Now().Add(time.Hour)
	//设置结构体
	claims := &userclaim{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirtime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   "user token",
		},
	}
	//给结构体产生一个token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//加密
	tokenstr, err := token.SignedString(key)
	if err != nil {
		fmt.Println("加密token出错", err)
	}
	return tokenstr
}

// AuthPublish 对用户投稿信息健全
func AuthPublish(c context.Context, ctx *app.RequestContext) {
	tokenstr := ctx.FormValue("token")
	fmt.Println("鉴权中间件token", tokenstr)
	if len(tokenstr) == 0 {
		fmt.Println("没有token")
		ctx.Abort()
		ctx.JSON(200, pojo.Response{
			StatusCode: -1,
			StatusMsg:  "Unauthorized",
		})
		return
	}
	//解析token得到结构体并检查是否有效
	token, muserclaim, err := ParseToken(string(tokenstr))
	if err != nil {
		fmt.Println("token解析出错")
		ctx.Abort()
		ctx.JSON(200, pojo.Response{
			StatusCode: -1,
			StatusMsg:  "Unauthorized",
		})
		return
	}
	if !token.Valid {
		//token无效,继续请求
		fmt.Println("token失效")
		ctx.Abort()
		ctx.JSON(200, pojo.Response{
			StatusCode: -1,
			StatusMsg:  "会话过期，请重新登录",
		})
		return
	} else {
		//token有效向上下文中加入
		fmt.Println("token里的userid", muserclaim.UserId)
		userid := strconv.Itoa(int(muserclaim.UserId))
		ctx.Set("userid", userid)
		ctx.Next(c)
	}

	return
}
