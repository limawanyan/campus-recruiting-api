package interceptor

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/proposal"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/redis"
	"campus-recruiting-api/internal/services/user"
	"go.uber.org/zap"
)

type Interceptor interface {
	// CheckLogin 验证是否登录
	CheckLogin(ctx core.Context) (info proposal.SessionUserInfo, err core.BusinessError)

	// GetLoginInfo 获取登录用户信息,未登录也不会阻止请求
	GetLoginInfo(ctx core.Context) (info proposal.SessionUserInfo, err core.BusinessError)

	// i 为了避免被其他包实现
	i()
}

type interceptor struct {
	logger      *zap.Logger
	cache       redis.Repo
	db          mysql.Repo
	userService user.Service
}

func New(logger *zap.Logger, cache redis.Repo, db mysql.Repo) Interceptor {
	return &interceptor{
		logger:      logger,
		cache:       cache,
		db:          db,
		userService: user.New(db, cache),
	}
}

func (i *interceptor) i() {}
