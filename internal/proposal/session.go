package proposal

import "encoding/json"

// SessionUserInfo 当前用户会话信息
type SessionUserInfo struct {
	UserID       int32  `json:"user_id"`       // 用户ID
	UserName     string `json:"user_name"`     // 用户名
	HeadPortrait string `json:"head_portrait"` // 头像
	Sex          int32  `json:"sex"`           // 性别
}

// Marshal 序列化到JSON
func (user *SessionUserInfo) Marshal() (jsonRaw []byte) {
	jsonRaw, _ = json.Marshal(user)
	return
}
