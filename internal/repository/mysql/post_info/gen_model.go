package post_info

import "time"

// PostInfo 平台帖子(面经/面题)
//go:generate gormgen -structs PostInfo -input .
type PostInfo struct {
	Id             int32     // 记录ID
	Title          string    // 标题
	Content        string    // 内容
	ParentType     int32     // 父级板块类型（面经/面题）
	SubType        int32     // 子板块类型
	CreatedAt      time.Time `gorm:"time"` // 创建时间
	UpdatedAt      time.Time `gorm:"time"` // 更新时间
	FromUserId     int32     // 发帖用户ID
	SortWeight     int64     // 热门排序权重（发帖时间戳+点赞数量*N）
	LikeNum        int32     // 点赞数量
	State          int32     // 状态，是否可用
	CommentNum     int32     // 评论数量
	BrowseNum      int32     // 浏览量
	IsDeleted      int32     // 是否逻辑删除（0 否 1是）
	CommentUpdated time.Time `gorm:"time"` // 最新评论时间
}
