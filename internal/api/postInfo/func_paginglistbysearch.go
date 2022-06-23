package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/postInfo"
	"net/http"
)

type pagingListBySearchRequest struct {
	Page       int    `form:"page"`       // 页码
	PageSize   int    `form:"pageSize"`   // 一页几条
	KeyWord    string `form:"keyWord"`    // 搜索关键字
	ParentType int32  `form:"parentType"` // 父级板块类型
	SubType    int32  `form:"subType"`    // 子版块类型
	SortType   int32  `form:"sortType"`   // 排序方式 最热 新回复 新发布
}

type pagingListBySearchResponse struct {
	List       []postInfo.PostListData `json:"list"`
	Pagination struct {
		Total        int64 `json:"total"`
		CurrentPage  int   `json:"currentPage"`
		PerPageCount int   `json:"perPageCount"`
	} `json:"pagination"`
}

// PagingListBySearch 模糊搜索分页获取面经/面题，从数据库获取，不限制最多为前20页
// @Summary 模糊搜索分页获取面经/面题，从数据库获取，不限制最多为前20页
// @Description 模糊搜索分页获取面经/面题，从数据库获取，不限制最多为前20页
// @Tags API.postInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body pagingListBySearchRequest true "请求信息"
// @Success 200 {object} pagingListBySearchResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/search [get]
func (h *handler) PagingListBySearch() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(pagingListBySearchRequest)
		res := new(pagingListBySearchResponse)
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

		searchData := new(postInfo.SearchData)
		searchData.Page = page
		searchData.PageSize = pageSize
		searchData.KeyWord = req.KeyWord
		searchData.ParentType = req.ParentType
		searchData.SubType = req.SubType
		searchData.SortType = req.SortType

		resListData, total, err := h.postService.PageListBySearch(ctx, searchData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.PostListBySearchError,
				code.Text(code.PostListBySearchError)).WithError(err),
			)
			return
		}

		res.List = resListData
		res.Pagination.Total = total
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page
		ctx.Payload(res)
	}
}
