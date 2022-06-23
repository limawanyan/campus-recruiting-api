package postInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/like_info"
)

func (s *service) ExistLike(ctx core.Context, fromUserId, postId, postType int32) (exist bool, err error) {
	postLikeQb := like_info.NewQueryBuilder()
	count, err := postLikeQb.
		WhereType(mysql.EqualPredicate, postType).
		WhereTopicId(mysql.EqualPredicate, postId).
		WhereFromUid(mysql.EqualPredicate, fromUserId).
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
