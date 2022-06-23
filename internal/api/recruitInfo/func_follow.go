package recruitInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
)

type followRequest struct {
	CompanyId int32 `json:"companyId" binding:"required"` // 公司ID
}

type followResponse struct {
	IsFollow bool `json:"isFollow"` // 关注
}

// Follow 关注或取消关注
// @Summary 关注或取消关注
// @Description 关注或取消关注
// @Tags API.recruitInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body followRequest true "请求信息"
// @Success 200 {object} followResponse
// @Failure 400 {object} code.Failure
// @Router /api/recruitInfo/follow [post]
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
		// 关注公司是否合法
		if exist, err := h.recruitInfoService.ExistCompany(ctx, req.CompanyId); err != nil || !exist {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		isFollow, err := h.recruitInfoService.Follow(ctx, req.CompanyId, ctx.SessionUserInfo().UserID) // 待完善
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.FollowCompanyError,
				code.Text(code.FollowCompanyError)).WithError(err))
			return
		}
		res.IsFollow = isFollow
		ctx.Payload(res)
	}
}
