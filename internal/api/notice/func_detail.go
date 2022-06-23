package notice

import (
	"campus-recruiting-api/internal/pkg/core"
)

type detailRequest struct{}

type detailResponse struct{}

// Detail 获取通知详情，如果通知为未读状态则设置成已读
// @Summary 获取通知详情，如果通知为未读状态则设置成已读
// @Description 获取通知详情，如果通知为未读状态则设置成已读
// @Tags API.notice
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/notice/detail [get]
func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
