package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/postInfo"
	"net/http"
)

type pagingListByIdsRequest struct {
	Ids []int32 `json:"ids" binding:"required"` // 要获取的数据ID集合
}

type pagingListByIdsResponse struct {
	Data []postInfo.PostListData `json:"data"`
}

// PagingListByIds 根据ID集合获取分页数据列表
// @Summary 根据ID集合获取分页数据列表
// @Description 根据ID集合获取分页数据列表
// @Tags API.postInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body pagingListByIdsRequest true "请求信息"
// @Success 200 {object} pagingListByIdsResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/list [get]
func (h *handler) PagingListByIds() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(pagingListByIdsRequest)
		res := new(pagingListByIdsResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		resListData, err := h.postService.ListByIds(ctx, req.Ids)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.GetPostListByIdsError,
				code.Text(code.GetPostListByIdsError)).WithError(err),
			)
			return
		}

		res.Data = resListData
		ctx.Payload(res)
	}
}
