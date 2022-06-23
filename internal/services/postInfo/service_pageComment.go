package postInfo

import "campus-recruiting-api/internal/pkg/core"

type GetPageCommentData struct {
	PostId   int32 `json:"postId" binding:"required"` // 帖子ID
	Page     int   `json:"page"`                      // 第几页
	PageSize int   `json:"pageSize"`                  // 每页显示条数
}

type PageCommentList struct {
	ID         int32             `json:"id"`
	ParentId   int32             `json:"parentId"`
	FromUserId int32             `json:"fromUserId"`
	ToUserId   int32             `json:"toUserId"`
	FromName   string            `json:"fromName"`
	FromHead   string            `json:"fromHead"`
	ToName     string            `json:"toName"`
	Content    string            `json:"content"`
	LikeNum    int32             `json:"likeNum"`
	Children   []PageCommentList `json:"children"`
}

func (s *service) PageCommentList(ctx core.Context, pageData *GetPageCommentData) (data []*PageCommentList, total int64, err error) {

	return
}
