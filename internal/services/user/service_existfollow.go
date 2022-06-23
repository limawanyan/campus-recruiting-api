package user

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user_follow"
)

func (s *service) ExistFollow(ctx core.Context, fromUserId, toUserId int32) (exist bool, err error) {
	userFollowQb := user_follow.NewQueryBuilder()
	count, err := userFollowQb.
		WhereType(mysql.EqualPredicate, int32(enum.User)).
		WhereTopicId(mysql.EqualPredicate, toUserId).
		WhereFromUserId(mysql.EqualPredicate, fromUserId).
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
