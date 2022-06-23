package user

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user"
)

func (s *service) ExistUser(ctx core.Context, userId int32) (exist bool, err error) {
	userQb := user.NewQueryBuilder()
	count, err := userQb.WhereId(mysql.EqualPredicate, userId).
		WhereState(mysql.EqualPredicate, int32(enum.Usable)).
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
