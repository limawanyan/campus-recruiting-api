package salaryInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
)

type voteRequest struct {
	ID        int32 `json:"id"`        // 帖子ID
	IsSupport bool  `json:"isSupport"` // 是否可信
}

type voteResponse struct {
	IsRepeat bool `json:"isRepeat"` // 是否重复
}

// Vote 投爆料可信或不可信（不可重复投）
// @Summary 投爆料可信或不可信（不可重复投）
// @Description 投爆料可信或不可信（不可重复投）
// @Tags API.salaryInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body voteRequest true "请求信息"
// @Success 200 {object} voteResponse
// @Failure 400 {object} code.Failure
// @Router /api/salaryInfo/vote [post]
func (h *handler) Vote() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(voteRequest)
		res := new(voteResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}

		isRepeat, err := h.salaryInfoService.Vote(ctx, req.ID, ctx.SessionUserInfo().UserID, req.IsSupport) // 待完善
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SalaryVoteError,
				code.Text(code.SalaryVoteError)).WithError(err))
			return
		}

		res.IsRepeat = isRepeat
		ctx.Payload(res)
	}
}
