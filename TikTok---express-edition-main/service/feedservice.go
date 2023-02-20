package service

import (
	"douyin/DAO"
	"douyin/pojo"
	"fmt"
	"log"
	"time"
)

func FeedList(latestTime time.Time) (result []pojo.Video, nextTime int64) {
	videos, err := DAO.FileDao.GetAllVideoWithPublishTime(latestTime)
	if err != nil {
		fmt.Println("获取视频流数据失败", err)
		return nil, -1
	}
	for idx, val := range videos {
		fmt.Println("idx:", idx)
		fmt.Println("val:", val)
	}
	result = make([]pojo.Video, 0, len(videos))
	nextTime, err = AssembleFeedVideos(&result, &videos)
	if err != nil {
		log.Println("方法copyVideos出错", err)
	}
	for idx, val := range result {
		fmt.Println("idx:", idx)
		fmt.Println("val:", val)
	}

	return result, nextTime
}

func AssembleFeedVideos(result *[]pojo.Video, data *[]DAO.TableVideo) (int64, error) {
	var timeMin time.Time = time.Now()
	for _, temp := range *data {
		var video pojo.Video
		//将video进行组装，添加想要的信息,插入从数据库中查到的数据
		//目前将comment_count,favorite_count以及is_favourite先固定化
		video.Id = temp.Id
		video.UserId = temp.AuthorId
		video.PlayUrl = temp.PlayUrl
		video.CoverUrl = temp.CoverUrl
		video.PublishTime = temp.PublishTime
		video.Title = temp.Title
		video.Author = *DAO.UserDAO.GetUserById(video.UserId)
		//固定化部分
		video.CommentCount = 1
		video.FavoriteCount = 1
		video.IsFavorite = true
		*result = append(*result, video)
		//获取最早时刻
		if video.PublishTime.Before(timeMin) {
			timeMin = video.PublishTime
		}
	}
	//作时间戳处理
	next_time := timeMin.Unix()
	return next_time, nil
}
