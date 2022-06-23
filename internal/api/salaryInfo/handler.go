package salaryInfo

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/redis"
	"campus-recruiting-api/internal/services/salaryInfo"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// PagingList 分页获取爆料信息
	// @Tags API.salaryInfo
	// @Router /api/salaryInfo [get]
	PagingList() core.HandlerFunc

	// Vote 投爆料可信或不可信（不可重复投）
	// @Tags API.salaryInfo
	// @Router /api/salaryInfo/vote [post]
	Vote() core.HandlerFunc

	// Detail 根据ID获取爆料详情
	// @Tags API.salaryInfo
	// @Router /api/salaryInfo/{id} [get]
	Detail() core.HandlerFunc

	// Create 薪资爆料
	// @Tags API.salaryInfo
	// @Router /api/salaryInfo/create [post]
	Create() core.HandlerFunc
}

type handler struct {
	logger            *zap.Logger
	db                mysql.Repo
	cache             redis.Repo
	salaryInfoService salaryInfo.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:            logger,
		cache:             cache,
		salaryInfoService: salaryInfo.New(db, cache),
	}
}

func (h *handler) i() {}
