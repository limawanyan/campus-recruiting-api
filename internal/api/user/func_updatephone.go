package user

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
)

type updatePhoneRequest struct{}

type updatePhoneResponse struct{}

// UpdatePhone 更新电话 暂不需要，通过微信小程序获取用户手机号
// @Summary 更新电话
// @Description 更新电话
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updatePhoneRequest true "请求信息"
// @Success 200 {object} updatePhoneResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/update_phone [post]
func (h *handler) UpdatePhone() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(updatePhoneRequest)
		res := new(updatePhoneResponse)
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
