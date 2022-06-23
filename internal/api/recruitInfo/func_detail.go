package recruitInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/recruitInfo"
	"net/http"
)

type detailRequest struct {
	CompanyId int32 `uri:"id" binding:"required"`
}

type detailResponse struct {
	Data recruitInfo.RecruitDetailDate `json:"data"`
}

// Detail 根据公司ID获取校招详细信息
// @Summary 根据公司ID获取校招详细信息
// @Description 根据公司ID获取校招详细信息
// @Tags API.recruitInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/recruitInfo/{id} [get]
func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		data, err := h.recruitInfoService.Detail(ctx, req.CompanyId)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.RecruitInfoDetailError,
				code.Text(code.RecruitInfoDetailError)).WithError(err))
			return
		}
		data.IsFollow, _ = h.recruitInfoService.ExistFollow(ctx, ctx.SessionUserInfo().UserID, req.CompanyId)
		res.Data = *data
		ctx.Payload(res)
	}
}
