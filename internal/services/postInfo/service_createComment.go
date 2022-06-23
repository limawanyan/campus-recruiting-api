package postInfo

import "campus-recruiting-api/internal/pkg/core"

type CreateCommentData struct {
	PostId     int32  // 评论帖子ID
	Type       int32  // 帖子类型 面经/面题
	ParentId   int32  // 父级评论ID
	ToUserId   int32  // 被评论人ID
	ToUserName string // 被评论人昵称
	ToUserHead string // 被评论人头像
	Content    string // 评论内容
	UserId     int32  // 评论人ID
	UserName   string // 评论人昵称
	UserHead   string // 评论人头像
}

func (s *service) CreateComment(ctx core.Context, createData *CreateCommentData) (id int32, err error) {
	return
}
