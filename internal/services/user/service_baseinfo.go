package user

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user"
)

type BaseInfoData struct {
	ID              int32  `json:"id"`
	Name            string `json:"name"`            // 用户名
	Sex             int32  `json:"sex"`             // 性别
	Head            string `json:"head"`            // 头像
	IsFollow        bool   `json:"isFollow"`        // 是否关注 当用户登录时获取
	FollowNum       int64  `json:"followNum"`       // 关注数量
	FansNum         int64  `json:"fansNum"`         // 粉丝数量
	GraduationYear  int    `json:"graduationYear"`  // 毕业年份
	Autograph       string `json:"autograph"`       // 签名
	IntentionalWork string `json:"intentionalWork"` // 意向工作
}

func (s *service) BaseInfo(ctx core.Context, userId int32) (data *BaseInfoData, err error) {
	userInfoQb := user.NewQueryBuilder()
	userInfo, err := userInfoQb.WhereIsDeleted(mysql.EqualPredicate, int32(enum.NotDeleted)).
		WhereId(mysql.EqualPredicate, userId).First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}
	data = new(BaseInfoData)
	data.Name = userInfo.Nickname
	data.Sex = userInfo.Sex
	data.Head = userInfo.HeadPortrait
	data.GraduationYear = userInfo.GraduationYear.Year()
	data.Autograph = userInfo.Autograph
	data.IntentionalWork = userInfo.IntentionalWork
	return
}
