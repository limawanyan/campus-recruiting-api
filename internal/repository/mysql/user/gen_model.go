package user

import "time"

// User 平台用户
//go:generate gormgen -structs User -input .
type User struct {
	Id              int32     // 数据唯一标识
	CreatedAt       time.Time `gorm:"time"` // 创建时间
	UpdatedAt       time.Time `gorm:"time"` // 更新时间
	WxOpenid        string    // 微信小程序授权openid
	Nickname        string    // 用户昵称
	Phone           string    // 手机号/账号
	Password        string    // 密码
	Sex             int32     // 用户性别
	HeadPortrait    string    // 用户头像
	GraduationYear  time.Time `gorm:"time"` // 毕业年份
	IntentionalWork string    // 意向工作
	Autograph       string    // 签名/简介
	Email           string    // 邮箱
	State           int32     // 状态，是否可用
	IsDeleted       int32     // 是否逻辑删除（0 否 1是）
}
