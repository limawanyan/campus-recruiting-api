package salaryInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/salaryInfo"
	"net/http"
)

type createRequest struct {
	CompanyName string `json:"companyName" binding:"required"` // 公司名称
	PostName    string `json:"postName" binding:"required"`    // 岗位名称
	CityName    string `json:"cityName" binding:"required"`    // 城市名称
	Description string `json:"description" binding:"required"` // 薪资描述
	Education   int32  `json:"education" binding:"required"`   // 学历
	Industry    int32  `json:"industry" binding:"required"`    // 行业
	Remarks     string `json:"remarks"`                        // 备注
}

type createResponse struct {
	Id int32 `json:"id"`
}

// Create 薪资爆料
// @Summary 薪资爆料
// @Description 薪资爆料
// @Tags API.salaryInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/salaryInfo/create [post]
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		if _, err := enum.EducationType(req.Education).ToStr(); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		if _, err := enum.IndustryType(req.Industry).ToStr(); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}

		createDate := new(salaryInfo.CreateSalaryInfoData)
		createDate.Description = req.Description
		createDate.CompanyName = req.CompanyName
		createDate.PostName = req.PostName
		createDate.CityName = req.CityName
		createDate.Remarks = req.Remarks
		createDate.Education = req.Education
		createDate.Industry = req.Industry
		createDate.FromUserId = ctx.SessionUserInfo().UserID
		createDate.FuzzyQuery = req.CompanyName + "_" + req.PostName + "_" + req.CityName

		id, err := h.salaryInfoService.Create(ctx, createDate)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SalaryInfoCreateError,
				code.Text(code.SalaryInfoCreateError)).WithError(err),
			)
			return
		}

		res.Id = id
		ctx.Payload(res)
	}
}
