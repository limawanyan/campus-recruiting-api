package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/postInfo"
	"net/http"
)

type getEditDetailRequest struct {
	PostId int32 `json:"postId" binding:"required"` // 帖子ID
}

type getEditDetailResponse struct {
	Data postInfo.EditDetailData `json:"data"`
}

// GetEditDetail 根据帖子ID获取帖子详细信息（标题、类别、内容）
// @Summary 根据帖子ID获取帖子详细信息（标题、类别、内容）
// @Description 根据帖子ID获取帖子详细信息（标题、类别、内容）
// @Tags API.postInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getEditDetailRequest true "请求信息"
// @Success 200 {object} getEditDetailResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/edit [get]
func (h *handler) GetEditDetail() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(getEditDetailRequest)
		res := new(getEditDetailResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		data, err := h.postService.EditDetail(ctx, req.PostId)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.GetPostEditDetailError,
				code.Text(code.GetPostEditDetailError)).WithError(err))
			return
		}
		res.Data = *data
		ctx.Payload(res)
	}
}
