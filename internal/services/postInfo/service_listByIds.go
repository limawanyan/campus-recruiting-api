package postInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/post_info"
	"campus-recruiting-api/internal/repository/mysql/user"
)

type PostListData struct {
	ID         int32    `json:"id"`         // 帖子ID
	Title      string   `json:"title"`      // 帖子标题
	Abstract   string   `json:"content"`    // 帖子摘要
	BrowseNum  int32    `json:"browseNum"`  // 浏览量
	CommentNum int32    `json:"commentNum"` // 评论数量
	LikeNum    int32    `json:"likeNum"`    // 评论数量
	User       UserData `json:"user"`       // 发布用户信息
}

func (s *service) ListByIds(ctx core.Context, Ids []int32) (data []PostListData, err error) {
	postInfoQb := post_info.NewQueryBuilder()
	postList, err := postInfoQb.WhereIdIn(Ids).
		WhereIsDeleted(mysql.EqualPredicate, int32(enum.NotDeleted)).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}
	data = make([]PostListData, len(postList))
	userIds := make([]int32, len(postList))
	for i, v := range postList {
		data[i].ID = v.Id
		data[i].Title = v.Title
		data[i].Abstract = string([]rune(v.Content)[:50])
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
