package postInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/post_info"
	"campus-recruiting-api/internal/repository/mysql/user"
	"regexp"
)

type SearchData struct {
	Page       int    // 页码
	PageSize   int    // 一页几条
	KeyWord    string // 搜索关键字
	ParentType int32  // 父级板块类型
	SubType    int32  // 子版块类型
	SortType   int32  // 排序方式
}

func (s *service) PageListBySearch(ctx core.Context, searchData *SearchData) (data []PostListData, total int64, err error) {
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
		WhereSubType(mysql.EqualPredicate, searchData.SubType)

	if searchData.KeyWord != "" {
		PostInfoQb.WhereTitle(mysql.LikePredicate, "%"+searchData.KeyWord+"%")
	}
	if enum.PostSortType(searchData.SortType) == enum.NewReply {
		PostInfoQb.OrderByCommentUpdated(false)
	} else if enum.PostSortType(searchData.SortType) == enum.NewRelease {
		PostInfoQb.OrderByCreatedAt(false)
	} else {
		// 最热
		PostInfoQb.OrderBySortWeight(false)
	}

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
		re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
		content := re.ReplaceAllString(v.Content, "\n")
		data[i].Abstract = content
		if len([]rune(content)) > 200 {
			data[i].Abstract = string([]rune(content)[:200])
		}
		data[i].BrowseNum = v.BrowseNum
		data[i].CommentNum = v.CommentNum
		data[i].LikeNum = v.LikeNum
		userIds[i] = v.FromUserId
	}

	// 获取作者信息
	userInfoQb := user.NewQueryBuilder()
	userList, err := userInfoQb.WhereIdIn(userIds).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}
	userMap := map[int32]*user.User{}
	for _, v := range userList {
		userMap[v.Id] = v
	}
	for i, v := range postList {
		data[i].User.ID = userMap[v.FromUserId].Id
		data[i].User.Name = userMap[v.FromUserId].Nickname
		data[i].User.IntentionalWork = userMap[v.FromUserId].IntentionalWork
		data[i].User.GraduationYear = userMap[v.FromUserId].GraduationYear.Year()
		data[i].User.Head = userMap[v.FromUserId].HeadPortrait
	}

	return
}
