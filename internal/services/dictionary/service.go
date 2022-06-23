package dictionary

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/redis"
	"go.uber.org/zap"
)

type Service interface {
	i()

	// SubItem 根据Code获取相应子项
	SubItem(ctx core.Context, code string) (data []SubItemData, err error)

	// IsSubItem 根据Code和value判断该Code是否存在值为value的子项
	IsSubItem(ctx core.Context, code string, value int32) (bool, error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}
