package video

import (
	"context"
	"douyin/pojo"
	"douyin/service"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"time"
)

func Feed(c context.Context, ctx *app.RequestContext) {
	LatestTime := ctx.Query("latest_time")
	//fmt.Println(latestTime)
	var videos []pojo.Video
	var nextTime int64
	if LatestTime == "" {
		loc, _ := time.LoadLocation("Local")
		latestTime, err := time.ParseInLocation("2006-01-02 15:04:05", LatestTime, loc)
		if err != nil {
			fmt.Println("id类型转换出错", err)
			return
		}
		videos, nextTime = service.FeedList(latestTime)
	} else {
		videos, nextTime = service.FeedList(time.Now())
	}

	ctx.JSON(200, pojo.FeedResponse{
		Response:  pojo.Response{StatusMsg: "", StatusCode: 0},
		NextTime:  nextTime,
		VideoList: videos,
	})

}
