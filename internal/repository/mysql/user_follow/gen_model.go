package user_follow

import "time"

// UserFollow 关注信息
//go:generate gormgen -structs UserFollow -input .
type UserFollow struct {
	Id         int32     // 数据唯一标识
	CreatedAt  time.Time `gorm:"time"` // 创建时间
	UpdatedAt  time.Time `gorm:"time"` // 更新时间
	Type       int32     // 关注类型（用户、校招公司...）
	TopicId    int32     // 主题ID
	FromUserId int32     // 收藏人
}
