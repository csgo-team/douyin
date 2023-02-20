package pojo

type UserResponse struct {
	Response
	User User `json:"user"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	StatusCode int32  `form:"Status_code" json:"Status_code" query:"Status_code"`
	StatusMsg  string `form:"Status_msg" json:"Status_msg" query:"Status_msg"`
	User       User   `json:"user"`
}

// FeedResponse 视频流接口的response
type FeedResponse struct {
	Response
	NextTime  int64   `json:"next_time"`
	VideoList []Video `json:"video_list"`
}

// PublishListResponse 发布列表的response
type PublishListResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"` //可为null
}

// IsLikeResponse 点赞Response
type IsLikeResponse struct {
	Response
}

// 评论操作的response
type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// 评论列表的response
type CommontListResponse struct {
	Response
	CommontList []Commentlist `json:"comment_list,omitempty"` //可为null
}
