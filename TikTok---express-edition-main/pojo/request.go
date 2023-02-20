package pojo

type IsLikeRequest struct {
	Token      string // 获取用户信息
	VideoId    string // 视屏id
	ActionType string // 1-点赞，2-取消点赞
}
