package postInfo

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/star_info"
)

func (s *service) Star(ctx core.Context, postId, starType, userId int32) (isStar bool, err error) {
	starQb := star_info.NewQueryBuilder()
	count, err := starQb.
		WhereType(mysql.EqualPredicate, starType).
		WhereFromUid(mysql.EqualPredicate, userId).
		WhereTopicId(mysql.EqualPredicate, postId).
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return false, err
	}

	// 取消收藏
	if count != 0 {
		if err = starQb.Delete(s.db.GetDbR().WithContext(ctx.RequestContext())); err != nil {
			return false, err
		}
		return
	}

	// 收藏
	isStar = true
	follow := star_info.NewModel()
	follow.Type = starType
	follow.TopicId = postId
	follow.FromUid = userId
	_, err = follow.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}
	return
}
