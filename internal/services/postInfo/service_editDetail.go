package postInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/post_info"
)

type EditDetailData struct {
	ID      int32  `json:"id"`      // 帖子ID
	Title   string `json:"title"`   // 帖子标题
	Content string `json:"content"` // 帖子内容
	SubType int32  `json:"subType"` // 所属栏目
}

func (s *service) EditDetail(ctx core.Context, postId int32) (editData *EditDetailData, err error) {
	postInfoQb := post_info.NewQueryBuilder()
	postInfo, err := postInfoQb.WhereId(mysql.EqualPredicate, postId).
		WhereIsDeleted(mysql.EqualPredicate, int32(enum.NotDeleted)).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}
	editData = new(EditDetailData)
	editData.ID = postInfo.Id
	editData.Title = postInfo.Title
	editData.Content = postInfo.Content
	editData.SubType = postInfo.SubType
	return
}
