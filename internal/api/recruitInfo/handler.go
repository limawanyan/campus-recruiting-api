package recruitInfo

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/redis"
	"campus-recruiting-api/internal/services/recruitInfo"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// PagingList 分页获取校招公司信息(如果用户选择了是否关注，则要从用户关注列表中查询)
	// @Tags API.recruitInfo
	// @Router /api/recruitInfo [get]
	PagingList() core.HandlerFunc

	// Detail 根据公司ID获取校招详细信息
	// @Tags API.recruitInfo
	// @Router /api/recruitInfo/{id} [get]
	Detail() core.HandlerFunc

	// Follow 关注或取消关注
	// @Tags API.recruitInfo
	// @Router /api/recruitInfo/follow [post]
	Follow() core.HandlerFunc
}

type handler struct {
	logger             *zap.Logger
	cache              redis.Repo
	recruitInfoService recruitInfo.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:             logger,
		cache:              cache,
		recruitInfoService: recruitInfo.New(db, cache),
	}
}

func (h *handler) i() {}
