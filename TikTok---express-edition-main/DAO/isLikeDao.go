package DAO

import (
	"douyin/Init/DB"
	"douyin/pojo"
	"strconv"
)

type videoDao struct {
	pojo.VideoLike
}

var VideoDao videoDao

func (v *videoDao) Add2Redis(videoId string, userIds []string) {
	dbClient := DB.GetDB()
	videoIdNum, _ := strconv.ParseInt(videoId, 10, 64)
	for userId := range userIds {
		// 从数组中获取每一个userID创建新用户并插入数据库
		videoL := pojo.VideoLike{
			VideoId: videoIdNum,
			UserId:  int64(userId),
			IsLike:  1,
		}
		// 存入数据库
		dbClient.Table("video_like").Create(&videoL)
	}

}
