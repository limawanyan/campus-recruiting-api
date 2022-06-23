package user

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user"
)

func (s *service) GetUserByWxOpenId(ctx core.Context, openId string) (data *user.User, err error) {
	userQb := user.NewQueryBuilder()
	data, err = userQb.WhereWxOpenid(mysql.EqualPredicate, openId).
		WhereState(mysql.EqualPredicate, int32(enum.Usable)).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	return
}
