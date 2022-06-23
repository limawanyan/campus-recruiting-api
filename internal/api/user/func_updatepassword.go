package user

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
)

type updatePasswordRequest struct{}

type updatePasswordResponse struct{}

// UpdatePassword 更新密码 暂不需要，通过微信小程序授权登录
// @Summary 更新密码
// @Description 更新密码
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updatePasswordRequest true "请求信息"
// @Success 200 {object} updatePasswordResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/update_password [post]
func (h *handler) UpdatePassword() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(updatePasswordRequest)
		res := new(updatePasswordResponse)
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
