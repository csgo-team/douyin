package DAO

import (
	"douyin/Init/DB"
	"time"
)

func (f *fileDao) GetAllVideoWithPublishTime(latestTime time.Time) ([]TableVideo, error) {
	db := DB.GetDB()
	var data []TableVideo
	result := db.Table("video").Where("created_at<?", latestTime).Order("created_at desc").Limit(5).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
