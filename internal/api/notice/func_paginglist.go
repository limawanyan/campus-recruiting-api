package notice

import (
	"campus-recruiting-api/internal/pkg/core"
)

type pagingListRequest struct{}

type pagingListResponse struct{}

// PagingList 根据通知类型分页获取通知列表
// @Summary 根据通知类型分页获取通知列表
// @Description 根据通知类型分页获取通知列表
// @Tags API.notice
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body pagingListRequest true "请求信息"
// @Success 200 {object} pagingListResponse
// @Failure 400 {object} code.Failure
// @Router /api/notice/list [get]
func (h *handler) PagingList() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
