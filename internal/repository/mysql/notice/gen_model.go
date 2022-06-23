package notice

import "time"

// Notice 通知
//go:generate gormgen -structs Notice -input .
type Notice struct {
	Id                  int32     // 记录ID
	CreatedAt           time.Time `gorm:"time"` // 创建时间
	UpdatedAt           time.Time `gorm:"time"` // 更新时间
	EventType           int32     // 事件类型（点赞、评论、回复）
	ContentType         int32     // 内容类型（面题、面经、面经评论、面题评论）
	Type                int32     // 通知类型（用户、系统）
	TriggerUserId       int32     // 触发人ID
	TriggerUserNickname string    // 触发人昵称
	TriggerUserPhoto    string    // 触发人头像
	SubjectId           int32     // 主题ID(面题、面经ID)
	ParentCommentId     int32     // 根评论ID
	CommentId           int32     // 评论ID
	TargetId            int32     // 接收人ID
	Title               string    // 标题
	Content             string    // 内容
	Uri                 string    // 内容链接
	IsRead              int32     // 是否已读(0 未读  1 已读)
}
