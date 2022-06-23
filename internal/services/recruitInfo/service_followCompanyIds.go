package recruitInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user_follow"
)

func (s *service) FollowCompanyIds(ctx core.Context, userId int32) (list map[int32]struct{}, err error) {
	follows, err := user_follow.NewQueryBuilder().
		WhereType(mysql.EqualPredicate, int32(enum.Company)).
		WhereFromUserId(mysql.EqualPredicate, userId).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	list = make(map[int32]struct{}, len(follows))
	for _, v := range follows {
		list[v.TopicId] = struct{}{}
	}
	return
}
