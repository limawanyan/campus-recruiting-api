package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
)

type likePostRequest struct {
	PostId int32 `json:"postId" binding:"required"`         // 帖子ID
	Type   int32 `json:"type" binding:"required,gt=0,lt=3"` // 点赞类型
}

type likePostResponse struct {
	IsLike bool `json:"isLike"` // 是否点赞
}

// LikePost 点赞帖子
// @Summary 点赞帖子
// @Description 点赞帖子
// @Tags API.postInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body likePostRequest true "请求信息"
// @Success 200 {object} likePostResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/like [post]
func (h *handler) LikePost() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(likePostRequest)
		res := new(likePostResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		// 验证帖子是否合法
		if exist, err := h.postService.ExistPost(ctx, req.PostId); err != nil || !exist {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		isLike, err := h.postService.Like(ctx, req.PostId, req.Type, ctx.SessionUserInfo().UserID)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.LikePostError,
				code.Text(code.LikePostError)).WithError(err))
			return
		}
		res.IsLike = isLike
		ctx.Payload(res)
	}
}
