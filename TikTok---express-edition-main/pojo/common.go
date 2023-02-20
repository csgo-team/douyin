package pojo

import (
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64     `gorm:"column:id" json:"id,omitempty"`
	Author        User      `json:"author,omitempty"`
	UserId        int64     `gorm:"column:userId"` //为了匹配数据库字段设置
	PlayUrl       string    `gorm:"column:play_url" json:"play_url,omitempty"`
	CoverUrl      string    `gorm:"column:cover_url" json:"cover_url,omitempty"`
	FavoriteCount int64     `gorm:"column:favourite_count" json:"favorite_count,omitempty"`
	CommentCount  int64     `gorm:"column:comment_count" json:"comment_count,omitempty"`
	IsFavorite    bool      `gorm:"column:is_favorite" json:"is_favorite,omitempty"`
	PublishTime   time.Time `gorm:"column:created_at"`
	Title         string    `gorm:"title" json:"title"`
}

type Comment struct {
	Id          int64  `gorm:"column:id" json:"id"`
	VideoId     int64  `gorm:"column:video_id"json:"-"` //一对多，视频对评论
	User        User   `json:"column:user" gorm:"-"`
	Content     string `gorm:"column:content" json:"content"`
	CreateDate  string `gorm:"column:create_date"json:"create_date"`
	commentId   string `gorm:"column:id"json:"comment_id"`
	actionType  string `gorm:"column:action_type"json:"action_type"`
	commentText string `gorm:"column:content"json:"comment_text"`
}

type Commentlist struct {
	Id         int64  `gorm:"column:id" json:"id"`
	User       User   `json:"user" gorm:"-"`
	Content    string `gorm:"column:content" json:"content"`
	CreateDate string `json:"create_date" gorm:"-"`
}

type User struct {
	Id            int64  `gorm:"column:id" json:"id,omitempty"`
	Name          string `gorm:"column:userName" json:"name,omitempty"`
	UserAccount   string `gorm:"column:userAccount" json:"user_account"`
	UserAvatar    string `gorm:"column:userAvatar" json:"user_avatar"`
	Gender        bool   `gorm:"column:gender" json:"gender"`
	Password      string `gorm:"column:userPassword" json:"password"`
	UserRole      int    `gorm:"column:userRole" json:"user_role"`
	AccessKey     string `gorm:"column:accessKey"`
	SecretKey     string `gorm:"column:secretKey"`
	FollowCount   int64  `gorm:"column:follow_count" json:"follow_count,omitempty"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	IsDelete      int8   `gorm:"isDelete" json:"is_delete"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type VideoLike struct {
	Id         int64     `gorm:"id" json:"id"`
	VideoId    int64     `gorm:"video_id" json:"video_Id"`
	UserId     int64     `gorm:"user_id" json:"user_Id"`
	IsLike     int8      `gorm:"isLike" json:"is_like"`
	CreateTime time.Time `gorm:"createTime" json:"create_time"`
}
