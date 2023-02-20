package service

import (
	"douyin/DAO"
	"douyin/pojo"
)

func GetCommentbyID(videoId int64) *pojo.Comment {
	return DAO.CommentDao.QueryCommentListByVideoId(videoId)
}
