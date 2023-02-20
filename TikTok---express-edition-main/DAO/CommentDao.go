package DAO

import (
	"douyin/Init/DB"
	"douyin/pojo"
	"errors"
	"gorm.io/gorm"
)

type CommentDAO struct {
	pojo.Comment
}

var (
	CommentDao CommentDAO
)

func NewCommentDAO() *CommentDAO {
	return &CommentDao
}

func (c *CommentDAO) AddCommentAndUpdateCount(comment *pojo.Comment) error {
	db := DB.GetDB()
	if comment == nil {
		return errors.New("AddCommentAndUpdateCount comment空指针")
	}
	//执行事务
	return db.Transaction(func(tx *gorm.DB) error {
		//添加评论数据
		if err := tx.Create(comment).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		//增加count
		if err := tx.Exec("UPDATE comments v SET v.comment_count = v.comment_count+1 WHERE v.id=?", comment.VideoId).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
}

// 删除
func (c *CommentDAO) DeleteCommentAndUpdateCountById(commentId, videoId int64) error {
	db := DB.GetDB()
	//执行事务
	return db.Transaction(func(tx *gorm.DB) error {

		//删除评论
		if err := tx.Exec("DELETE FROM comments WHERE id = ?", commentId).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		//减少count
		if err := tx.Exec("UPDATE comments v SET v.comment_count = v.comment_count-1 WHERE v.id=? AND v.comment_count>0", videoId).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

// 查找
func (c *CommentDAO) QueryCommentById(id int64, comment *pojo.Comment) error {
	db := DB.GetDB()
	if comment == nil {
		return errors.New("QueryCommentById comment 空指针")
	}
	return db.Where("id=?", id).First(comment).Error
}

// CommentList
func (c *CommentDAO) QueryCommentListByVideoId(videoId int64) *pojo.Comment {
	db := DB.GetDB()
	db.Model(&pojo.Comment{}).Where("video_id=?", videoId).Find(&c.Comment)
	return &c.Comment
}
