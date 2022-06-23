package postInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/post_info"
)

type EditData struct {
	ID         int32
	Title      string // 标题
	Content    string // 内容
	ParentType int32  // 父级板块类型
	SubType    int32  // 子版块类型
	FromUserId int32  // 编辑人
	SortWeight int64  // 热门排序权重
}

func (s *service) Edit(ctx core.Context, data *EditData) error {
	var err error
	if data.ID != 0 {
		// 编辑
		postInfoQb := post_info.NewQueryBuilder()
		_, err = postInfoQb.WhereId(mysql.EqualPredicate, data.ID).
			WhereIsDeleted(mysql.EqualPredicate, int32(enum.NotDeleted)).
			WhereFromUserId(mysql.EqualPredicate, data.FromUserId).
			First(s.db.GetDbW().WithContext(ctx.RequestContext()))
		if err != nil {
			return err
		}
		err = postInfoQb.Updates(s.db.GetDbW().WithContext(ctx.RequestContext()), map[string]interface{}{
			"title":    data.Title,
			"content":  data.Content,
			"sub_type": data.SubType,
		})
		if err != nil {
			return err
		}
	} else {
		// 创建
		postInfo := post_info.NewModel()
		postInfo.Title = data.Title
		postInfo.Content = data.Content
		postInfo.ParentType = data.ParentType
		postInfo.SubType = data.SubType
		postInfo.FromUserId = data.FromUserId
		postInfo.SortWeight = data.SortWeight
		postInfo.State = int32(enum.Usable)
		_, err = postInfo.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
		if err != nil {
			return err
		}
	}
	return err
}
