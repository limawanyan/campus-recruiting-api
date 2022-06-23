package user

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
)

type registerRequest struct {
	WxOpenId string `json:"wxOpenId"`
}

type registerResponse struct{}

// Register 注册 暂不使用，需要适用H5端时使用
// @Summary 注册
// @Description 注册
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body registerRequest true "请求信息"
// @Success 200 {object} registerResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/Register [post]
func (h *handler) Register() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(registerRequest)
		res := new(registerResponse)
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
