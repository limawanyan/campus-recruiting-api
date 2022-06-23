package postInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/post_info"
)

type SearchByUserData struct {
	Page       int   // 第几页
	PageSize   int   // 一页多少条
	ParentType int32 // 板块类型 面经/面题
	UserId     int32 // 当前登录用户
}

func (s *service) PageListByUser(ctx core.Context, searchData *SearchByUserData) (data []PostListData, total int64, err error) {

	page := searchData.Page
	if page <= 0 {
		page = 1
	}
	pageSize := searchData.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	PostInfoQb := post_info.NewQueryBuilder()
	PostInfoQb.WhereIsDeleted(mysql.EqualPredicate, int32(enum.NotDeleted)).
		WhereParentType(mysql.EqualPredicate, searchData.ParentType).
		WhereFromUserId(mysql.EqualPredicate, searchData.UserId)

	total, err = PostInfoQb.Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}

	postList, err := PostInfoQb.Limit(pageSize).Offset(offset).QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}

	data = make([]PostListData, len(postList))
	for i, v := range postList {
		data[i].ID = v.Id
		data[i].Title = v.Title
		data[i].BrowseNum = v.BrowseNum
		data[i].CommentNum = v.CommentNum
		data[i].LikeNum = v.LikeNum
		if len(v.Content) > 80 {
			data[i].Abstract = string([]rune(v.Content)[:80])
		} else {
			data[i].Abstract = v.Content
		}
	}

	return
}
