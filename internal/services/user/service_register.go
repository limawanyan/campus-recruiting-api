package user

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql/user"
)

type RegisterUserData struct {
	WxOpenId     string // 微信OpenId
	NickName     string // 昵称
	Sex          int32  // 性别
	HeadPortrait string // 头像
}

func (s *service) Register(ctx core.Context, data *RegisterUserData) (id int32, err error) {
	user := user.NewModel()
	user.WxOpenid = data.WxOpenId
	user.Sex = data.Sex
	user.HeadPortrait = data.HeadPortrait
	user.Nickname = data.NickName
	id, err = user.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	return
}
