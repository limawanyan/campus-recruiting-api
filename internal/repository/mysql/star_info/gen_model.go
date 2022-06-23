package star_info

import "time"

// StarInfo 收藏信息
//go:generate gormgen -structs StarInfo -input .
type StarInfo struct {
	Id        int32     // 记录ID
	CreatedAt time.Time `gorm:"time"` // 创建时间
	UpdatedAt time.Time `gorm:"time"` // 更新时间
	FromUid   int32     // 收藏用户
	Type      int32     // 收藏类型（面经、面题）
	TopicId   int32     // 收藏主题
}
