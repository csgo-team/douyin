package service

import (
	"context"
	"crypto/sha256"
	"douyin/DAO"
	"douyin/Init/Redis"
	"douyin/config"
	"douyin/middleware"
	"douyin/pojo"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"time"
)

// Iflogin 登陆
func Iflogin(ctx context.Context, username, password string) (loginResponse pojo.UserLoginResponse) {
	user := DAO.UserDAO.Getuserbyname(username)
	if user == nil || user.Password != Encoder(password) {
		//nil说明数据库中没有该用户 ====== =
		if user == nil || user.Password != Encoder(password) {
			// nil说明数据库中没有该用户 >>>>>>> 8d89065bb80e9478c862a67e741a22967dda7be7
			loginResponse = pojo.UserLoginResponse{
				Response: pojo.Response{StatusMsg: "用户名或者密码错误", StatusCode: -1},
				UserId:   -1,
				Token:    "",
			}
		} else {
			//登录签发一个token
			tokenstr := middleware.SetToken(*user)
			// 将登陆后的用户设置到redis中 key是useridfo:id  value是用户自身
			redisClient := Redis.InitRedis()
			key := config.USER_INFO_KEY + strconv.Itoa(int(user.Id))
			tokenKey := config.TOKEN_KEY + strconv.Itoa(int(user.Id))
			// 存入redis中
			if err := redisClient.Set(ctx, key, user, time.Minute*10).Err(); err != nil {
				return pojo.UserLoginResponse{
					Response: pojo.Response{StatusMsg: "redis存入失败", StatusCode: -1},
				}
			}
			// 将token也存进redis中，下次获取先从token中获取
			redisClient.Set(ctx, tokenKey, tokenstr, time.Minute*10)

			loginResponse = pojo.UserLoginResponse{
				Response: pojo.Response{StatusMsg: "Login Success", StatusCode: 200},
				UserId:   user.Id,
				Token:    tokenstr,
			}
		}
		return
	}
	return
}

// IfRegister 注册服务
func IfRegister(username, password string) (registerResponse pojo.UserLoginResponse) {
	//查看数据库中是否已存在该用户
	user := DAO.UserDAO.Getuserbyname(username)
	if len(user.UserAccount) != 0 {
		registerResponse = pojo.UserLoginResponse{
			Response: pojo.Response{StatusMsg: "用户名已存在", StatusCode: -1},
			UserId:   -1,
		}
		return
	}
	//组装新用户
	user = &pojo.User{
		Name:          "user" + username,
		UserAccount:   username,
		Password:      Encoder(password),
		UserRole:      0,
		FollowerCount: 0,
		FollowCount:   0}
	//调用dao来添加用户
	user = DAO.UserDAO.CreateUser(*user)
	if user == nil {
		//从数据库拿不到刚刚注册的用户，说明注册失败
		registerResponse = pojo.UserLoginResponse{
			Response: pojo.Response{StatusMsg: "注册失败", StatusCode: -1},
			UserId:   -1,
		}
	} else {
		//获取签发token
		tokenstr := middleware.SetToken(*user)
		fmt.Println("token", tokenstr)
		registerResponse = pojo.UserLoginResponse{
			Response: pojo.Response{StatusMsg: "注册成功", StatusCode: 0},
			UserId:   user.Id,
			Token:    tokenstr,
		}
	}
	return
}
func GetUserById(ctx context.Context, id int64) pojo.UserInfoResponse {
	// 从redis中获取并返回就行
	redis := Redis.InitRedis()
	idStr := strconv.Itoa(int(id))
	key := config.USER_INFO_KEY + idStr
	userRedis := &pojo.User{}
	err := redis.Get(ctx, key).Scan(userRedis)
	if err != nil {
		return pojo.UserInfoResponse{
			StatusCode: 1,
			StatusMsg:  "redis获取失败",
		}
	}

	return pojo.UserInfoResponse{
		StatusCode: 200,
		StatusMsg:  "成功获取用户信息",
		User:       *userRedis,
	}
}

// Encoder
/***
给密码加密的函数
*/
func Encoder(password string) (passwordEncoded string) {
	w := sha256.New()
	io.WriteString(w, password)
	bw := w.Sum(nil) //w.Sum(nil)将w的hash转成[]byte格式
	passwordEncoded = hex.EncodeToString(bw)
	return
}
