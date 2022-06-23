package user

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user_follow"
)

func (s *service) FansNum(ctx core.Context, userId int32) (count int64, err error) {
	followQb := user_follow.NewQueryBuilder()
	count, err = followQb.WhereType(mysql.EqualPredicate, int32(enum.User)).
		WhereTopicId(mysql.EqualPredicate, userId).Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	return
}
