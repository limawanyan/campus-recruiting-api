package salaryInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/salaryInfo"
	"net/http"
)

type detailRequest struct {
	ID int32 `uri:"id" binding:"required"` // 公司ID
}

type detailResponse struct {
	Data salaryInfo.SalaryDetailData `json:"data"`
}

// Detail 根据ID获取爆料详情
// @Summary 根据ID获取爆料详情
// @Description 根据ID获取爆料详情
// @Tags API.salaryInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/salaryInfo/{id} [get]
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
		data, err := h.salaryInfoService.Detail(ctx, req.ID)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SalaryInfoDetailError,
				code.Text(code.SalaryInfoDetailError)).WithError(err))
			return
		}
		res.Data = *data
		ctx.Payload(res)
	}
}
