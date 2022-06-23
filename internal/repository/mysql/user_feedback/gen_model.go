package user_feedback

import "time"

// UserFeedback 用户反馈
//go:generate gormgen -structs UserFeedback -input .
type UserFeedback struct {
	Id        int32     //
	UserId    int32     // 反馈用户
	Content   string    // 反馈内容
	Tel       string    // 联系电话
	Type      int32     // 反馈类型(优化建议、Bug反馈)
	State     int32     // 状态（0 待处理、1已处理）
	Reply     string    //
	CreatedAt time.Time `gorm:"time"` //
	UpdatedAt time.Time `gorm:"time"` //
}
