package user

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
)

type followRequest struct {
	UserId int32 `json:"userId" binding:"required"` // 被关注用户ID
}

type followResponse struct {
	IsFollow bool `json:"isFollow"` // 是否关注
}

// Follow 关注/取消关注用户
// @Summary 关注/取消关注用户
// @Description 关注/取消关注用户
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body followRequest true "请求信息"
// @Success 200 {object} followResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/follow [post]
func (h *handler) Follow() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(followRequest)
		res := new(followResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		// 不能关注自己
		if ctx.SessionUserInfo().UserID == req.UserId {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.FocusSelfError,
				code.Text(code.FocusSelfError)))
			return
		}
		// 用户是否合法
		if exist, err := h.userService.ExistUser(ctx, req.UserId); err != nil || !exist {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		isFollow, err := h.userService.Follow(ctx, req.UserId, 1) // 待完善
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.FollowUserError,
				code.Text(code.FollowUserError)).WithError(err))
			return
		}
		res.IsFollow = isFollow
		ctx.Payload(res)
	}
}
