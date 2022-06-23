package router

import (
	"campus-recruiting-api/internal/router/interceptor"
	"fmt"

	"campus-recruiting-api/configs"
	"campus-recruiting-api/internal/alert"
	"campus-recruiting-api/internal/metrics"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/redis"
	"campus-recruiting-api/pkg/errors"

	"go.uber.org/zap"
)

type resource struct {
	mux          core.Mux
	logger       *zap.Logger
	db           mysql.Repo
	cache        redis.Repo
	interceptors interceptor.Interceptor
}

type Server struct {
	Mux   core.Mux
	Db    mysql.Repo
	Cache redis.Repo
}

func NewHTTPServer(logger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}
	r := new(resource)
	r.logger = logger

	// 浏览器打开地址
	openBrowserUri := configs.ProjectDomain + configs.ProjectPort

	// 初始化DB
	dbRepo, err := mysql.New()
	if err != nil {
		logger.Fatal("new db err", zap.Error(err))
	}
	r.db = dbRepo

	// 初始化Cache
	cacheRepo, err := redis.New()
	if err != nil {
		fmt.Println("Redis初始化失败，请配置Redis。")
		logger.Fatal("new cache err", zap.Error(err))
	}
	r.cache = cacheRepo

	mux, err := core.New(
		logger,
		core.WithEnableOpenBrowser(openBrowserUri),
		core.WithEnableCors(),
		core.WithEnableRate(),
		core.WithAlertNotify(alert.NotifyHandler(logger)),
		core.WithRecordMetrics(metrics.RecordHandler(logger)),
	)

	if err != nil {
		panic(err)
	}

	r.mux = mux
	r.interceptors = interceptor.New(logger, r.cache, r.db)

	// 设置API路由
	setApiRouter(r)

	s := new(Server)
	s.Mux = mux
	s.Db = r.db
	s.Cache = r.cache

	return s, nil
}
