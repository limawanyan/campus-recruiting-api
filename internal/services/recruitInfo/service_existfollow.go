package recruitInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user_follow"
)

func (s *service) ExistFollow(ctx core.Context, userId, companyId int32) (exist bool, err error) {
	followQb := user_follow.NewQueryBuilder()
	count, err := followQb.WhereTopicId(mysql.EqualPredicate, companyId).
		WhereFromUserId(mysql.EqualPredicate, userId).
		WhereType(mysql.EqualPredicate, int32(enum.Company)).
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
