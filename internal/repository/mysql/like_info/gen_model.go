package like_info

import "time"

// LikeInfo 点赞信息表
//go:generate gormgen -structs LikeInfo -input .
type LikeInfo struct {
	Id        int32     // 记录ID
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
	TopicId   int32     // 主题ID
	Type      int32     // 点赞类型（讨论帖/面题/讨论帖评论/面题评论/薪资可信度）
	FromUid   int32     // 来自哪个用户
	State     int32     // 状态，点赞是否有效(0 无效 1 有效)
}
