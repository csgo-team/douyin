package service

import (
	"context"
	"douyin/DAO"
	"douyin/Init/Redis"
	"douyin/config"
	"douyin/pojo"
	"fmt"
	idworker "github.com/gitstliu/go-id-worker"
	uuid "github.com/satori/go.uuid"
	"log"
	"mime/multipart"
	"sync"
	"time"
)

func Publish(data *multipart.FileHeader, userId int64, title string) pojo.Response {
	file, err := data.Open()
	if err != nil {
		log.Println("视频文件打开错误", err)
		return pojo.Response{
			StatusCode: 404,
			StatusMsg:  "视频文件上传失败",
		}
	}
	//给视频生成唯一标识名称和id
	videoname, vid := GetUniqueId()
	//保存到云端
	err = DAO.FileDao.SaveVideo(file, videoname+config.PlayUrlSuffix)
	if err != nil {
		log.Println("保存文件到云失败", err)
		return pojo.Response{
			StatusCode: 404,
			StatusMsg:  "视频文件上传失败",
		}
	}
	//给封面图生成唯一标识名称，忽略它的id
	covername, _ := GetUniqueId()
	//把封面图保存到云端
	coverurl, err := DAO.FileDao.GetCoverUrl(covername, videoname)
	if err != nil {
		log.Println("获取截图出错", err)
		return pojo.Response{
			StatusCode: 404,
			StatusMsg:  "视频文件上传失败",
		}
	}
	//拼接视频访问url
	videourl := config.PlayUrlPrefix + videoname + config.PlayUrlSuffix
	//组转video
	video := *NewVideoObject(vid, videourl, coverurl, title, userId)
	fmt.Println("PlayUrl打印", video.PlayUrl)
	//将video记录放进数据库
	DAO.FileDao.SaveVideoIntoDataBase(video)
	return pojo.Response{
		StatusCode: 200,
		StatusMsg:  "上传成功",
	}
}

func PublishList(id int64) pojo.PublishListResponse {
	videos, err := DAO.FileDao.GetAllVideoWithUserId(id)
	if err != nil {
		fmt.Println("获取视频数据失败", err)
		return pojo.PublishListResponse{Response: pojo.Response{
			StatusCode: 1, StatusMsg: "获取失败",
		},
		}
	}
	for idx, val := range videos {
		fmt.Println("idx:", idx)
		fmt.Println("val:", val)
	}
	result := make([]pojo.Video, 0, len(videos))
	err = AssembleVideos(&result, &videos, id)
	if err != nil {
		log.Println("方法copyVideos出错", err)
	}
	res := pojo.PublishListResponse{Response: pojo.Response{
		StatusCode: 0, StatusMsg: "用户列表已获取",
	}, VideoList: result,
	}
	return res
}

// GetUniqueId 生成视频的唯一辨识id
func GetUniqueId() (videoname string, videoid int64) {
	//创建
	id := uuid.NewV4().String()
	fmt.Println("uuid:", id)
	year, month, day := time.Now().Date()
	//将uuid和时间拼接生成唯一标识名称
	vidname := fmt.Sprintf("%v%v%v%v", year, month, day, id)
	fmt.Println("视频id：", vidname)
	//生成存入数据库的id标识
	currWoker := &idworker.IdWorker{}
	currWoker.InitIdWorker(1000, 1)
	newId, newIdErr := currWoker.NextId()
	if newIdErr != nil {
		log.Println("获取视频数据库id错误", newIdErr)
	}
	return vidname, newId
}

// Id          int64     `gorm:"column:id" json:"id"`
// AuthorId    int64     `gorm:"column:userId"`
// PlayUrl     string    `gorm:"column:play_url" json:"play_url"`
// CoverUrl    string    `gorm:"column:cover_url json:"cover_url"`
// PublishTime time.Time `gorm:"column:created_at"`
// Title
// 将tablevideo转换成video
func AssembleVideos(result *[]pojo.Video, data *[]DAO.TableVideo, userId int64) error {
	for _, temp := range *data {
		var video pojo.Video
		//将video进行组装，添加想要的信息,插入从数据库中查到的数据
		//目前将comment_count,favorite_count以及is_favourite先固定化
		video.Id = temp.Id
		video.PlayUrl = temp.PlayUrl
		video.CoverUrl = temp.CoverUrl
		video.PublishTime = temp.PublishTime
		video.Title = temp.Title
		video.Author = *DAO.UserDAO.GetUserById(userId)
		//固定化部分
		video.CommentCount = 1
		video.FavoriteCount = 1
		video.IsFavorite = true
		*result = append(*result, video)
	}
	return nil
}

func NewVideoObject(vid int64, playurl, coverurl, title string, userid int64) *DAO.TableVideo {
	return &DAO.TableVideo{
		Id:          vid,
		AuthorId:    userid,
		PlayUrl:     playurl,
		CoverUrl:    coverurl,
		Title:       title,
		PublishTime: time.Now(),
	}
}

// DoLike 点赞
func DoLike(c context.Context, videoId string, userID string, isLike int64) *pojo.IsLikeResponse {
	redisClient := Redis.InitRedis()
	// 判断传入的参数为0还是1
	// 当isLike为1时，点赞状态
	key := config.VIDEO_LIKE_KEY + videoId

	// 单独开启线程执行
	var wg sync.WaitGroup
	go func() {
		// 当点赞数量超过3000条时，写入数据库
		members := redisClient.SMembers(c, key)
		if len(members.Val()) >= 3000 {
			defer wg.Done()
			// 调用Dao层的数据
			DAO.VideoDao.Add2Redis(videoId, members.Val())
		}
	}()
	wg.Wait()

	if isLike == 1 {
		add := redisClient.SAdd(c, key, userID)
		if add.Val() == 1 {
			return &pojo.IsLikeResponse{
				Response: pojo.Response{
					StatusCode: 0,
					StatusMsg:  "点赞成功",
				},
			}
		} else {
			return &pojo.IsLikeResponse{
				Response: pojo.Response{
					StatusCode: -1,
					StatusMsg:  "重复点赞",
				},
			}
		}
	}

	// 当isLike为0时，取消点赞
	if isLike == 0 {
		// 先判断是否点过赞
		isMember := redisClient.SIsMember(c, key, videoId)
		if !isMember.Val() {
			return &pojo.IsLikeResponse{
				Response: pojo.Response{
					StatusCode: -1,
					StatusMsg:  "请先登陆",
				},
			}
		}
		// 从redis中删除
		rem := redisClient.SRem(c, key)
		if rem.Val() == 1 {
			return &pojo.IsLikeResponse{
				Response: pojo.Response{
					StatusCode: 0,
					StatusMsg:  "取消成功",
				},
			}
		} else {
			return &pojo.IsLikeResponse{
				Response: pojo.Response{
					StatusCode: -1,
					StatusMsg:  "点赞失效",
				},
			}
		}
	}
	return &pojo.IsLikeResponse{
		Response: pojo.Response{
			StatusCode: -1,
			StatusMsg:  "isLike参数错误",
		},
	}
}
