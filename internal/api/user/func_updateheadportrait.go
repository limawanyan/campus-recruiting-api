package user

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
)

type updateHeadPortraitRequest struct {
	Head string `json:"head" binding:"required"` // 新头像
}

type updateHeadPortraitResponse struct{}

// UpdateHeadPortrait 更新头像 通过微信授权获取新头像
// @Summary 更新头像
// @Description 更新头像
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateHeadPortraitRequest true "请求信息"
// @Success 200 {object} updateHeadPortraitResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/head_portrait [post]
func (h *handler) UpdateHeadPortrait() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(updateHeadPortraitRequest)
		res := new(updateHeadPortraitResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}

		ctx.Payload(res)
	}
}
