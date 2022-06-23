package postInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/post_info"
)

func (s *service) ExistPost(ctx core.Context, postId int32) (exist bool, err error) {
	postQb := post_info.NewQueryBuilder()
	count, err := postQb.WhereId(mysql.EqualPredicate, postId).
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
