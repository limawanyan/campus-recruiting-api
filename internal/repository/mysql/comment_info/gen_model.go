package comment_info

import "time"

// CommentInfo 评论信息
//go:generate gormgen -structs CommentInfo -input .
type CommentInfo struct {
	Id           int32     // 记录ID
	CreatedAt    time.Time `gorm:"time"` // 创建时间
	UpdatedAt    time.Time `gorm:"time"` // 更新时间
	TopicId      int32     // 主题ID
	Type         int32     // 评论类型（讨论帖/面题/薪资爆料）
	ParentId     int32     // 父级评论ID
	FromUserId   int32     // 评论人
	ToUserId     int32     // 被评论人
	FromNickname string    // 评论人昵称
	FromHead     string    // 评论人头像
	ToNickname   string    // 被评论人昵称
	Content      string    // 评论内容
	LikeNum      int32     // 点赞数量
	IsDeleted    int32     // 是否逻辑删除（0 否 1是）
}
