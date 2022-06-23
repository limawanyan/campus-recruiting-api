package dictionary

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/redis"
	"campus-recruiting-api/internal/services/dictionary"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// SubItem 根据字典标识获取字典子项
	// @Tags API.dictionary
	// @Router /api/dictionary/subItem/{code} [get]
	SubItem() core.HandlerFunc
}

type handler struct {
	logger            *zap.Logger
	cache             redis.Repo
	dictionaryService dictionary.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:            logger,
		cache:             cache,
		dictionaryService: dictionary.New(logger, db, cache),
	}
}

func (h *handler) i() {}
