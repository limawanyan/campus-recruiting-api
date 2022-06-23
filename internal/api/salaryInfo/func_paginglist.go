package salaryInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/salaryInfo"
	"net/http"
	"time"
)

type pagingListRequest struct {
	Page      int    `form:"page"`      // 第几页
	PageSize  int    `form:"pageSize"`  // 每页显示条数
	Industry  int32  `form:"industry"`  // 所属行业
	Education int32  `form:"education"` // 公司属性
	KeyWord   string `form:"keyWord"`   // 模糊搜索关键字
}

type salaryInfoData struct {
	ID          int32     `json:"id"`
	CompanyName string    `json:"company"`     // 公司名称
	PostName    string    `json:"postName"`    // 岗位名称
	Education   string    `json:"education"`   // 学历
	Reliability int32     `json:"reliability"` // 可信度
	Description string    `json:"description"` // 薪资描述
	CityName    string    `json:"cityName"`    // 城市
	CreatedAt   time.Time `json:"createdAt"`   // 发布时间
}

type pagingListResponse struct {
	List       []salaryInfoData `json:"list"`
	Pagination struct {
		Total        int64 `json:"total"`
		CurrentPage  int   `json:"currentPage"`
		PerPageCount int   `json:"perPageCount"`
	} `json:"pagination"`
}

// PagingList 分页获取爆料信息
// @Summary 分页获取爆料信息
// @Description 分页获取爆料信息
// @Tags API.salaryInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body pagingListRequest true "请求信息"
// @Success 200 {object} pagingListResponse
// @Failure 400 {object} code.Failure
// @Router /api/salaryInfo [get]
func (h *handler) PagingList() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(pagingListRequest)
		res := new(pagingListResponse)
		if err := ctx.ShouldBindQuery(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		page := req.Page
		if page <= 0 {
			page = 1
		}

		pageSize := req.PageSize
		if pageSize <= 0 {
			pageSize = 10
		}

		searchData := new(salaryInfo.SearchData)
		searchData.Page = page
		searchData.PageSize = pageSize
		searchData.Industry = req.Industry
		searchData.Education = req.Education
		searchData.KeyWord = req.KeyWord

		resListData, total, err := h.salaryInfoService.PageList(ctx, searchData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.SalaryInfoListError,
				code.Text(code.SalaryInfoListError)).WithError(err),
			)
			return
		}

		res.Pagination.Total = total
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page

		res.List = make([]salaryInfoData, len(resListData))
		for k, v := range resListData {
			res.List[k].ID = v.Id
			res.List[k].CompanyName = v.CompanyName
			res.List[k].PostName = v.PostName
			res.List[k].Education, _ = enum.EducationType(v.Education).ToStr()
			res.List[k].Reliability = v.Reliability
			res.List[k].Description = v.Description
			res.List[k].CityName = v.CityName
			res.List[k].CreatedAt = v.CreatedAt
		}

		ctx.Payload(res)
	}
}
