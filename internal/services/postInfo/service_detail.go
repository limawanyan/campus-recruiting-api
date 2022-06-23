package postInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/post_info"
	"campus-recruiting-api/internal/repository/mysql/user"
	"time"
)

type PostDetailData struct {
	ID         int32     `json:"id"`        // 帖子ID
	Title      string    `json:"title"`     // 帖子标题
	Content    string    `json:"content"`   // 帖子内容
	CreatedAt  time.Time `json:"createdAt"` // 发布时间
	User       UserData  `json:"user"`      // 发布用户信息
	IsStar     bool      `json:"isStar"`    // 是否收藏
	IsLike     bool      `json:"isLike"`    // 是否点赞
	ParentType int32     `json:"type"`      // 帖子类型
}

type UserData struct {
	ID              int32  `json:"id"`              // 用户ID
	Name            string `json:"name"`            // 用户名称
	IntentionalWork string `json:"intentionalWork"` // 意向工作
	GraduationYear  int    `json:"graduation"`      // 毕业年份
	Head            string `json:"head"`            // 头像
	IsFollow        bool   `json:"isFocus"`         // 是否关注
}

func (s *service) Detail(ctx core.Context, postId int32) (postDetail *PostDetailData, err error) {
	postInfoQb := post_info.NewQueryBuilder()
	postInfo, err := postInfoQb.WhereId(mysql.EqualPredicate, postId).
		WhereIsDeleted(mysql.EqualPredicate, int32(enum.NotDeleted)).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}
	postDetail = new(PostDetailData)
	postDetail.ID = postInfo.Id
	postDetail.Title = postInfo.Title
	postDetail.Content = postInfo.Content
	postDetail.CreatedAt = postInfo.CreatedAt
	postDetail.ParentType = postInfo.ParentType
	// 获取作者信息
	userInfoQb := user.NewQueryBuilder()
	user, err := userInfoQb.WhereId(mysql.EqualPredicate, postInfo.FromUserId).
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}
	postDetail.User.ID = user.Id
	postDetail.User.Name = user.Nickname
	postDetail.User.IntentionalWork = user.IntentionalWork
	postDetail.User.GraduationYear = user.GraduationYear.Year()
	postDetail.User.Head = user.HeadPortrait
	return
}
