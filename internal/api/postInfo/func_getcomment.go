package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/postInfo"
	"net/http"
)

type getCommentRequest struct {
	PostId   int32 `form:"postId" binding:"required"` // 帖子ID
	Page     int   `form:"page"`                      // 第几页
	PageSize int   `form:"pageSize"`                  // 每页显示条数
}

type getCommentResponse struct {
	Data       []*postInfo.PageCommentList `json:"data"`
	Pagination struct {
		Total        int64 `json:"total"`
		CurrentPage  int   `json:"currentPage"`
		PerPageCount int   `json:"perPageCount"`
	} `json:"pagination"`
}

// GetComment 分页获取帖子评论
// @Summary 获取帖子评论
// @Description 获取帖子评论
// @Tags API.postInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getCommentRequest true "请求信息"
// @Success 200 {object} getCommentResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/comment [get]
func (h *handler) GetComment() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(getCommentRequest)
		res := new(getCommentResponse)
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

		searchData := new(postInfo.GetPageCommentData)
		searchData.PostId = req.PostId
		searchData.Page = req.Page
		searchData.PageSize = req.PageSize

		resListData, total, err := h.postService.PageCommentList(ctx, searchData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.GetCommentListError,
				code.Text(code.GetCommentListError)).WithError(err),
			)
			return
		}

		res.Data = resListData
		res.Pagination.Total = total
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page
		ctx.Payload(res)
	}
}
