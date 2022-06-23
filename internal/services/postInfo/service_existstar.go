package postInfo

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/like_info"
)

func (s *service) ExistStar(ctx core.Context, fromUserId, postId, postType int32) (exist bool, err error) {
	postStarQb := like_info.NewQueryBuilder()
	count, err := postStarQb.
		WhereType(mysql.EqualPredicate, postType).
		WhereTopicId(mysql.EqualPredicate, postId).
		WhereFromUid(mysql.EqualPredicate, fromUserId).
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
