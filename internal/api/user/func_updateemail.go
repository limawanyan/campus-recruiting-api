package user

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
)

type updateEmailRequest struct{}

type updateEmailResponse struct{}

// UpdateEmail 更新邮箱 暂不需要，后期需要绑定邮箱操作再实现
// @Summary 更新邮箱
// @Description 更新邮箱
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateEmailRequest true "请求信息"
// @Success 200 {object} updateEmailResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/update_email [post]
func (h *handler) UpdateEmail() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(updateEmailRequest)
		res := new(updateEmailResponse)
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
