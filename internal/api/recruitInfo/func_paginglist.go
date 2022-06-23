package recruitInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/recruitInfo"
	"net/http"
)

type pagingListRequest struct {
	Page      int   `form:"page"`         // 第几页
	PageSize  int   `form:"pageSize"`     // 每页显示条数
	Industry  int32 `form:"industryType"` // 所属行业
	Attribute int32 `form:"attribute"`    // 公司属性
	IsFollow  bool  `form:"isFollow"`     // 用户是否关注
}

type recruitInfoData struct {
	ID              int32  `json:"id"`
	Name            string `json:"name"`            // 公司名称
	RecommendTime   string `json:"recommendTime"`   // 内推时间
	ApplyOnlineTime string `json:"applyOnlineTime"` // 网申时间
	WrittenExamTime string `json:"writtenExamTime"` // 笔试时间
	InterviewTime   string `json:"interviewTime"`   // 面试时间
	OfferTime       string `json:"offerTime"`       // Offer发放时间
	Industry        string `json:"industry"`        // 所属行业
	Logo            string `json:"logo"`            // Logo
	IsFollow        bool   `json:"isFollow"`        // 已关注
}

type pagingListResponse struct {
	List       []recruitInfoData `json:"list"`
	Pagination struct {
		Total        int64 `json:"total"`
		CurrentPage  int   `json:"currentPage"`
		PerPageCount int   `json:"perPageCount"`
	} `json:"pagination"`
}

// PagingList 分页获取校招公司信息(如果用户选择了是否关注，则要从用户关注列表中查询)
// @Summary 分页获取校招公司信息(如果用户选择了是否关注，则要从用户关注列表中查询)
// @Description 为了适配是否关注筛选条件，日程模块必须用户登录
// @Tags API.recruitInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body pagingListRequest true "请求信息"
// @Success 200 {object} pagingListResponse
// @Failure 400 {object} code.Failure
// @Router /api/recruitInfo [get]
func (h *handler) PagingList() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(pagingListRequest)
		res := new(pagingListResponse)
		if err := ctx.ShouldBindForm(req); err != nil {
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

		searchData := new(recruitInfo.SearchData)
		searchData.Page = page
		searchData.PageSize = pageSize
		searchData.Industry = req.Industry
		searchData.Attribute = req.Attribute
		searchData.IsFollow = req.IsFollow
		searchData.UserId = ctx.SessionUserInfo().UserID

		resListData, total, err := h.recruitInfoService.PageList(ctx, searchData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.RecruitInfoListError,
				code.Text(code.RecruitInfoListError)).WithError(err),
			)
			return
		}

		res.Pagination.Total = total
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page

		// 获取用户关注的所有公司
		followIds, err := h.recruitInfoService.FollowCompanyIds(ctx, ctx.SessionUserInfo().UserID)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.RecruitInfoListError,
				code.Text(code.RecruitInfoListError)).WithError(err),
			)
			return
		}

		res.List = make([]recruitInfoData, len(resListData))
		for k, v := range resListData {
			res.List[k].ID = v.Id
			res.List[k].Name = v.CompanyName
			res.List[k].RecommendTime = v.RecommendTime
			res.List[k].ApplyOnlineTime = v.ApplyOnlineTime
			res.List[k].WrittenExamTime = v.WrittenExamTime
			res.List[k].InterviewTime = v.InterviewTime
			res.List[k].OfferTime = v.OfferTime
			res.List[k].Industry = v.Industrys
			res.List[k].Logo = v.CompanyLogo
			if _, ok := followIds[v.Id]; ok {
				res.List[k].IsFollow = true
			}
		}

		ctx.Payload(res)
	}
}
