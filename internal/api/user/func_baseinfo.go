package user

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/user"
	"go.uber.org/zap"
	"net/http"
)

type baseInfoRequest struct {
	UserId int32 `json:"userId" binding:"required"`
}

type baseInfoResponse struct {
	Data user.BaseInfoData `json:"data"`
}

// BaseInfo 获取基础信息
// @Summary 获取基础信息
// @Description 获取基础信息
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body baseInfoRequest true "请求信息"
// @Success 200 {object} baseInfoResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/base_info [get]
func (h *handler) BaseInfo() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(baseInfoRequest)
		res := new(baseInfoResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		data, err := h.userService.BaseInfo(ctx, req.UserId)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.BaseUserInfoError,
				code.Text(code.BaseUserInfoError)).WithError(err))
			return
		}
		data.FansNum, err = h.userService.FansNum(ctx, data.ID)
		if err != nil {
			h.logger.Error("获取用户粉丝数量出错", zap.Error(err))
		}
		data.FollowNum, err = h.userService.FollowNum(ctx, data.ID)
		if err != nil {
			h.logger.Error("获取用户关注数量出错", zap.Error(err))
		}
		// 如果当前请求用户登录了则获取是否关注该用户
		// 如果是查看自己的主页，关注按钮隐藏
		data.IsFollow, err = h.userService.ExistFollow(ctx, 1, data.ID)
		if err != nil {
			h.logger.Error("获取用户是否关注数据出错", zap.Error(err))
		}

		res.Data = *data
		ctx.Payload(res)
	}
}
