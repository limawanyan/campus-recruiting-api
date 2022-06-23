package notice

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// PagingList 根据通知类型分页获取通知列表
	// @Tags API.notice
	// @Router /api/notice/list [get]
	PagingList() core.HandlerFunc

	// Detail 获取通知详情，如果通知为未读状态则设置成已读
	// @Tags API.notice
	// @Router /api/notice/detail [get]
	Detail() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
}

func New(logger *zap.Logger, db mysql.Repo) Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) i() {}
