package user

import (
	"campus-recruiting-api/internal/pkg/core"
)

type loginOutRequest struct {
}

type loginOutResponse struct{}

// LoginOut 登出
// @Summary 登出
// @Description 登出
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body loginOutRequest true "请求信息"
// @Success 200 {object} loginOutResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/loginOut [post]
func (h *handler) LoginOut() core.HandlerFunc {
	return func(ctx core.Context) {
		res := new(loginOutResponse)
		// 清除当前用户登录状态 token作废
		ctx.Payload(res)
	}
}
