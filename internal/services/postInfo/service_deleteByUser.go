package postInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/post_info"
)

func (s *service) DeleteByUser(ctx core.Context, postId, userId int32) error {
	postInfoQb := post_info.NewQueryBuilder()
	_, err := postInfoQb.WhereId(mysql.EqualPredicate, postId).
		WhereFromUserId(mysql.EqualPredicate, userId).
		WhereIsDeleted(mysql.EqualPredicate, int32(enum.NotDeleted)).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return err
	}
	err = postInfoQb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), map[string]interface{}{
		"is_deleted": 1,
	})
	return err
}
