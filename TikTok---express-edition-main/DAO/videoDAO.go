package DAO

import (
	"douyin/Init/DB"
	"douyin/config"
	"encoding/base64"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"log"
	"time"
)

type fileDao struct{}

var FileDao fileDao

// 对接video表的结构体
type TableVideo struct {
	Id          int64     `gorm:"column:id" json:"id"`
	AuthorId    int64     `gorm:"column:userId"`
	PlayUrl     string    `gorm:"column:play_url" json:"play_url"`
	CoverUrl    string    `gorm:"column:cover_url" json:"cover_url"`
	PublishTime time.Time `gorm:"column:created_at"`
	Title       string    `gorm:"title" json:"title"` //视频名，5.23添加
}

// 此函数使tablevideo结构体可以绑定到video表
func (TableVideo) TableName() string {
	return "video"
}

// 实现视频文件的简单文件上传
func (f *fileDao) SaveVideo(file io.Reader, filename string) (err error) {
	//获取OSSClient实例
	client, err := oss.New("oss-cn-guangzhou.aliyuncs.com", config.AccessKeyID, config.AccessKeyPassword)
	if err != nil {
		fmt.Println("获取OSSclient失败")
		log.Println(err)
		return
	}
	//获取存储空间
	bucket, err := client.Bucket(config.BucketName)

	if err != nil {
		log.Println("获取存储空间失败", err)
		return
	}
	//将filename扩展成带有目录层级的结构,这样才能存到存储空间的二级目录
	filename = "video/" + filename
	err = bucket.PutObject(filename, file)
	if err != nil {
		log.Println("文件上传到云失败", err)
		return
	}
	return nil
}

// 实现图片文件的简单文件上传
func (f *fileDao) SaveCover(file io.Reader, filename string) (err error) {
	//获取OSSClient实例
	client, err := oss.New("oss-cn-guangzhou.aliyuncs.com", config.AccessKeyID, config.AccessKeyPassword)
	if err != nil {
		fmt.Println("获取OSSclient失败")
		log.Println(err)
		return
	}
	//获取存储空间
	bucket, err := client.Bucket(config.BucketName)
	if err != nil {
		log.Println("获取存储空间失败", err)
		return
	}
	//将filename扩展成带有目录层级的结构,这样才能存到存储空间的二级目录
	filename = "cover/" + filename
	err = bucket.PutObject(filename, file)
	if err != nil {
		log.Println("文件上传到云失败", err)
		return
	}
	return nil
}

// 存进数据库
func (f *fileDao) SaveVideoIntoDataBase(video TableVideo) {
	db := DB.GetDB()
	db.Table("video").Create(&video)
	return
}

/*
**
获取封面的url
*/
func (f *fileDao) GetCoverUrl(CoverName, videopath string) (url string, err error) {
	//获取OSSClient实例
	client, err := oss.New("oss-cn-guangzhou.aliyuncs.com", config.AccessKeyID, config.AccessKeyPassword)
	if err != nil {
		fmt.Println("获取OSSclient失败")
		log.Println(err)
		return "", err
	}
	//获取存储空间
	bucket, err := client.Bucket(config.BucketName)
	if err != nil {
		log.Println("获取存储空间失败", err)
		return "", err
	}
	targetsource := "cover/" + CoverName + config.CoverUrlSuffix
	videosource := "video/" + videopath + config.PlayUrlSuffix
	style := "video/snapshot,t_1000,f_jpg,w_480,h_800"
	process := fmt.Sprintf("%s|sys/saveas,o_%v,b_%v", style, base64.URLEncoding.EncodeToString([]byte(targetsource)), base64.URLEncoding.EncodeToString([]byte(config.BucketName)))
	result, err := bucket.ProcessObject(videosource, process)
	if err != nil {
		log.Println("云端生成截图失败", err)
		return "", err
	}
	fmt.Println(result)
	return config.CoverUrlPrefix + CoverName + config.CoverUrlSuffix, err
}

func (f *fileDao) GetAllVideoWithUserId(id int64) ([]TableVideo, error) {
	db := DB.GetDB()
	var data []TableVideo
	result := db.Table("video").Where("userId=? && isDelete = 0", id).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}
