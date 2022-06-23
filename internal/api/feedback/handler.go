package feedback

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/redis"
	"campus-recruiting-api/internal/services/feedback"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 用户提交反馈
	// @Tags API.feedback
	// @Router /api/feedback/create [post]
	Create() core.HandlerFunc

	// List 用户获取反馈记录
	// @Tags API.feedback
	// @Router /api/feedback/list [get]
	List() core.HandlerFunc
}

type handler struct {
	logger          *zap.Logger
	cache           redis.Repo
	feedbackService feedback.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:          logger,
		cache:           cache,
		feedbackService: feedback.New(db, cache),
	}
}

func (h *handler) i() {}
