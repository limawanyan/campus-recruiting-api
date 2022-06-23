package user

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user_follow"
)

func (s *service) Follow(ctx core.Context, userId, fromUserId int32) (isFollow bool, err error) {
	followQb := user_follow.NewQueryBuilder()
	count, err := followQb.
		WhereType(mysql.EqualPredicate, int32(enum.User)).
		WhereFromUserId(mysql.EqualPredicate, fromUserId).
		WhereTopicId(mysql.EqualPredicate, userId).
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return false, err
	}

	// 取消关注
	if count != 0 {
		isFollow = false
		if err = followQb.Delete(s.db.GetDbR().WithContext(ctx.RequestContext())); err != nil {
			return false, err
		}
		return
	}

	// 关注
	isFollow = true
	follow := user_follow.NewModel()
	follow.Type = int32(enum.User)
	follow.TopicId = userId
	follow.FromUserId = fromUserId
	_, err = follow.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}
	return
}
