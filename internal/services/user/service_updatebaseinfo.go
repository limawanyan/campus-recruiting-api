package user

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user"
	"time"
)

type UpdateBaseInfoData struct {
	NikeName        string    // 昵称
	GraduationYear  time.Time // 毕业年份
	IntentionalWork string    // 意向岗位
	Autograph       string    // 签名
	UserId          int32     // 更新用户
}

func (s *service) UpdateBaseInfo(ctx core.Context, updateData *UpdateBaseInfoData) error {
	userInfoQb := user.NewQueryBuilder()
	_, err := userInfoQb.WhereIsDeleted(mysql.EqualPredicate, int32(enum.NotDeleted)).
		WhereId(mysql.EqualPredicate, updateData.UserId).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return err
	}
	err = userInfoQb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), map[string]interface{}{
		"nickname":         updateData.NikeName,
		"graduation_year":  updateData.GraduationYear,
		"intentional_work": updateData.IntentionalWork,
		"autograph":        updateData.Autograph,
	})
	return err
}
