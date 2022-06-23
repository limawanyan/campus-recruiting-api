package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
)

type starPostRequest struct {
	PostId int32 `json:"postId" binding:"required"`         // 帖子ID
	Type   int32 `json:"type" binding:"required,gt=0,lt=3"` // 收藏类型
}

type starPostResponse struct {
	IsStar bool `json:"isStar"` // 是否收藏
}

// StarPost 收藏帖子
// @Summary 收藏帖子
// @Description 收藏帖子
// @Tags API.postInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body starPostRequest true "请求信息"
// @Success 200 {object} starPostResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/star [post]
func (h *handler) StarPost() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(starPostRequest)
		res := new(starPostResponse)
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
		isStar, err := h.postService.Star(ctx, req.PostId, req.Type, ctx.SessionUserInfo().UserID)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.StarPostError,
				code.Text(code.StarPostError)).WithError(err))
			return
		}
		res.IsStar = isStar
		ctx.Payload(res)
	}
}
