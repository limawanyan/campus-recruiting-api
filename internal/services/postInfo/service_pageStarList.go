package postInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/post_info"
	"campus-recruiting-api/internal/repository/mysql/star_info"
)

func (s *service) PageStarList(ctx core.Context, searchData *SearchByUserData) (data []PostListData, total int64, err error) {
	page := searchData.Page
	if page <= 0 {
		page = 1
	}
	pageSize := searchData.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	starQb := star_info.NewQueryBuilder()
	starList, err := starQb.WhereType(mysql.EqualPredicate, searchData.ParentType).
		WhereFromUid(mysql.EqualPredicate, searchData.UserId).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}
	starIds := make([]int32, len(starList))
	for i, v := range starList {
		starIds[i] = v.TopicId
	}

	PostInfoQb := post_info.NewQueryBuilder()
	PostInfoQb.WhereIdIn(starIds).WhereIsDeleted(mysql.EqualPredicate, int32(enum.NotDeleted))
	total, err = PostInfoQb.Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}

	postList, err := PostInfoQb.Limit(pageSize).Offset(offset).QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}

	// 帖子信息
	data = make([]PostListData, len(postList))
	userIds := make([]int32, len(postList))
	for i, v := range postList {
		data[i].ID = v.Id
		data[i].Title = v.Title
		data[i].BrowseNum = v.BrowseNum
		data[i].CommentNum = v.CommentNum
		data[i].LikeNum = v.LikeNum
		userIds[i] = v.FromUserId
		if len(v.Content) > 80 {
			data[i].Abstract = string([]rune(v.Content)[:80])
		} else {
			data[i].Abstract = v.Content
		}
	}

	return
}
