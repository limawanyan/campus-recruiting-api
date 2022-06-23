package postInfo

import (
	"campus-recruiting-api/configs"
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/like_info"
)

func (s *service) Like(ctx core.Context, postId, likeType, userId int32) (isLike bool, err error) {
	likeQb := like_info.NewQueryBuilder()
	likeInfo, err := likeQb.
		WhereType(mysql.EqualPredicate, likeType).
		WhereFromUid(mysql.EqualPredicate, userId).
		WhereTopicId(mysql.EqualPredicate, postId).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	// 取消点赞 代码待优化 需要使用事务
	if likeInfo != nil {
		if likeInfo.State == int32(enum.Usable) {
			err = likeQb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), map[string]interface{}{
				"state": 0,
			})
		} else {
			err = likeQb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), map[string]interface{}{
				"state": 1,
			})
			isLike = true
		}
		if err != nil {
			return false, err
		}
		s.db.GetDbW().Exec("UPDATE post_info SET like_num = like_num-1,sort_weight = sort_weight-? WHERE id = ?", configs.HotValue, postId)
	} else {
		isLike = true
		like := like_info.NewModel()
		like.Type = likeType
		like.TopicId = postId
		like.FromUid = userId
		like.State = int32(enum.Usable)
		_, err = like.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
		if err != nil {
			return
		}
		s.db.GetDbW().Exec("UPDATE post_info SET like_num = like_num+1,sort_weight = sort_weight+? WHERE id = ?", configs.HotValue, postId)
	}
	return
}
