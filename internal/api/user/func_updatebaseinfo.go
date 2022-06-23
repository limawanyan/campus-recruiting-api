package user

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/user"
	"net/http"
	"time"
)

type updateBaseInfoRequest struct {
	NikeName        string    `json:"nikeName" binding:"required"`        // 昵称
	GraduationYear  time.Time `json:"graduationYear" binding:"required"`  // 毕业年份
	IntentionalWork string    `json:"intentionalWork" binding:"required"` // 意向岗位
	Autograph       string    `json:"autograph"`                          // 签名
	Email           string    `json:"email" binding:"email"`              // 邮箱
}

type updateBaseInfoResponse struct{}

// UpdateBaseInfo 更新基础信息
// @Summary 更新基础信息
// @Description 更新基础信息
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body updateBaseInfoRequest true "请求信息"
// @Success 200 {object} updateBaseInfoResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/update [post]
func (h *handler) UpdateBaseInfo() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(updateBaseInfoRequest)
		res := new(updateBaseInfoResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}

		updateDate := new(user.UpdateBaseInfoData)
		updateDate.NikeName = req.NikeName
		updateDate.GraduationYear = req.GraduationYear
		updateDate.IntentionalWork = req.IntentionalWork
		updateDate.Autograph = req.Autograph
		updateDate.UserId = 1 // 待完善

		err := h.userService.UpdateBaseInfo(ctx, updateDate)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.UpdateUserBaseInfoError,
				code.Text(code.UpdateUserBaseInfoError)).WithError(err),
			)
			return
		}

		ctx.Payload(res)
	}
}
