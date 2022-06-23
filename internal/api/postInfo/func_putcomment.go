package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/postInfo"
	"net/http"
)

type putCommentRequest struct {
	PostId     int32  `json:"postId" binding:"required"`  // 评论帖子ID
	Type       int32  `json:"type" binding:"required"`    // 帖子类型 面经/面题
	ParentId   int32  `json:"parentId"`                   // 父级评论ID
	ToUserId   int32  `json:"toUserId"`                   // 被评论人ID
	ToUserName string `json:"toUserName"`                 // 被评论人昵称
	ToUserHead string `json:"toUserHead"`                 // 被评论人头像
	Content    string `json:"content" binding:"required"` // 评论内容

}

type putCommentResponse struct {
	ID int32 `json:"id"`
}

// PutComment 发表评论
// @Summary 发表评论
// @Description 发表评论
// @Tags API.postInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body putCommentRequest true "请求信息"
// @Success 200 {object} putCommentResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/comment [post]
func (h *handler) PutComment() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(putCommentRequest)
		res := new(putCommentResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}

		createDate := new(postInfo.CreateCommentData)
		createDate.Type = req.Type
		createDate.ParentId = req.ParentId
		createDate.ToUserName = req.ToUserName
		createDate.ToUserHead = req.ToUserHead
		createDate.ToUserId = req.ToUserId
		createDate.Content = req.Content
		createDate.UserHead = ""
		createDate.UserName = "Admin" // 待完善
		createDate.UserId = ctx.SessionUserInfo().UserID

		id, err := h.postService.CreateComment(ctx, createDate)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.CreatePostCommentError,
				code.Text(code.CreatePostCommentError)).WithError(err),
			)
			return
		}

		res.ID = id
		ctx.Payload(res)

	}
}
