package feedback

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user_feedback"
)

func (s *service) List(ctx core.Context, userId int32) (listData []*user_feedback.UserFeedback, err error) {
	feedbackQb := user_feedback.NewQueryBuilder()
	feedbackQb.WhereUserId(mysql.EqualPredicate, userId)
	listData, err = feedbackQb.QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
