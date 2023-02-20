package comment

import (
	"context"
	"douyin/DAO"
	"douyin/pojo"
	"fmt"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func CommentAction(c context.Context, ctx *app.RequestContext) {
	video_id := ctx.Query("video_id")
	id, err := strconv.ParseInt(video_id, 10, 64)
	if err != nil {
		fmt.Println("视频id有误")
	}

	//获取用户评论
	comment_text := ctx.Query("comment_text")
	fmt.Println(comment_text)

	//解析token后获得的userId,在DAO找到该id的用户信息
	userid := ctx.GetString("userId")
	userId, err := strconv.ParseInt(userid, 10, 64)
	user := DAO.UserDAO.GetUserById(userId)
	fmt.Println(user.Name)

	// 1-发布评论，2-删除评论
	action_type := ctx.Query("action_type")
	if action_type == "1" {
		ctx.JSON(200, pojo.CommentActionResponse{
			Response: pojo.Response{StatusMsg: "评论成功", StatusCode: 0},
			Comment:  pojo.Comment{Id: id, Content: comment_text, CreateDate: time.Now().Format("2006-01-02 15:04:05"), User: pojo.User{Id: user.Id, Name: user.Name, FollowCount: user.FollowCount, FollowerCount: user.FollowerCount, IsFollow: user.IsFollow}},
		})

	} else if action_type == "2" {
		ctx.JSON(200, pojo.CommentActionResponse{
			Response: pojo.Response{StatusMsg: "删除成功", StatusCode: 0},
		})

	} else {
		ctx.JSON(200, pojo.CommentActionResponse{
			Response: pojo.Response{StatusMsg: "操作失败", StatusCode: -1},
		})

	}

}
