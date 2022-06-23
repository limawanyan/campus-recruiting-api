package postInfo

import (
	"campus-recruiting-api/configs"
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/post_info"
)

type GetIdsData struct {
	ParentType int32 // 父板块类型
	SubType    int32 // 子版块类型
	SortType   int32 // 排序方式 最热 新回复 新发布
}

func (s *service) GetListIds(ctx core.Context, searchData *GetIdsData) (data []int32, err error) {
	PostInfoQb := post_info.NewQueryBuilder()
	PostInfoQb.WhereIsDeleted(mysql.EqualPredicate, int32(enum.NotDeleted)).
		WhereParentType(mysql.EqualPredicate, searchData.ParentType).
		WhereSubType(mysql.EqualPredicate, searchData.SubType)
	if enum.PostSortType(searchData.SortType) == enum.NewReply {
		PostInfoQb.OrderByCommentUpdated(false)
	} else if enum.PostSortType(searchData.SortType) == enum.NewRelease {
		PostInfoQb.OrderByCreatedAt(false)
	} else {
		// 最热
		PostInfoQb.OrderBySortWeight(false)
	}
	postList, err := PostInfoQb.Limit(configs.IdsMaxPage * configs.PageSize).QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}
	data = make([]int32, len(postList))
	for i, v := range postList {
		data[i] = v.Id
	}
	return
}
