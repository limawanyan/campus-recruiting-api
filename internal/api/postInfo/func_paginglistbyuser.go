package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/postInfo"
	"net/http"
)

type pagingListByUserRequest struct {
	Page       int   `json:"page"`                          // 第几页
	PageSize   int   `json:"pageSize"`                      // 每页显示条数
	ParentType int32 `json:"parentType" binding:"required"` // 父板块类型
}

type pagingListByUserResponse struct {
	Data       []postInfo.PostListData `json:"data"`
	Pagination struct {
		Total        int64 `json:"total"`
		CurrentPage  int   `json:"currentPage"`
		PerPageCount int   `json:"perPageCount"`
	} `json:"pagination"`
}

// PagingListByUser 分页获取用户发布的面经/面题,按时间倒序
// @Summary 分页获取用户发布的面经/面题,按时间倒序
// @Description 分页获取用户发布的面经/面题,按时间倒序
// @Tags API.postInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body pagingListByUserRequest true "请求信息"
// @Success 200 {object} pagingListByUserResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/user [get]
func (h *handler) PagingListByUser() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(pagingListByUserRequest)
		res := new(pagingListByUserResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
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

		searchData := new(postInfo.SearchByUserData)
		searchData.Page = page
		searchData.PageSize = pageSize
		searchData.ParentType = req.ParentType
		searchData.UserId = ctx.SessionUserInfo().UserID

		resListData, total, err := h.postService.PageListByUser(ctx, searchData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.PagePostListByUserError,
				code.Text(code.PagePostListByUserError)).WithError(err),
			)
			return
		}

		res.Pagination.Total = total
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page
		res.Data = resListData
		ctx.Payload(res)
	}
}
