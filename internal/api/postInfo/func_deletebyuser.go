package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
)

type deleteByUserRequest struct {
	PostId int32 `json:"postId" binding:"required"` // 帖子ID
}

type deleteByUserResponse struct {
}

// DeleteByUser 用户删除自己发布的面经/面题
// @Summary 用户删除自己发布的面经/面题
// @Description 用户删除自己发布的面经/面题,执行逻辑删除模式
// @Tags API.postInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body deleteByUserRequest true "请求信息"
// @Success 200 {object} deleteByUserResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/delete [post]
func (h *handler) DeleteByUser() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(deleteByUserRequest)
		res := new(deleteByUserResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		err := h.postService.DeleteByUser(ctx, req.PostId, ctx.SessionUserInfo().UserID) // 待完善
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.DeletePostByUserError,
				code.Text(code.DeletePostByUserError)).WithError(err),
			)
			return
		}
		ctx.Payload(res)
	}
}
