package user

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/redis"
	"campus-recruiting-api/internal/services/user"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Login 登录
	// @Tags API.user
	// @Router /api/user/login [post]
	Login() core.HandlerFunc

	// LoginOut 登出
	// @Tags API.user
	// @Router /api/user/loginOut [post]
	LoginOut() core.HandlerFunc

	// Register 注册
	// @Tags API.user
	// @Router /api/user/Register [post]
	Register() core.HandlerFunc

	// UpdateHeadPortrait 更新头像
	// @Tags API.user
	// @Router /api/user/head_portrait [post]
	UpdateHeadPortrait() core.HandlerFunc

	// BaseInfo 获取基础信息
	// @Tags API.user
	// @Router /api/user/info [get]
	BaseInfo() core.HandlerFunc

	// UpdateBaseInfo 更新基础信息
	// @Tags API.user
	// @Router /api/user/update [post]
	UpdateBaseInfo() core.HandlerFunc

	// UpdatePassword 更新密码
	// @Tags API.user
	// @Router /api/user/updatePassword [post]
	UpdatePassword() core.HandlerFunc

	// UpdateEmail 更新邮箱
	// @Tags API.user
	// @Router /api/user/updateEmail [post]
	UpdateEmail() core.HandlerFunc

	// UpdatePhone 更新电话
	// @Tags API.user
	// @Router /api/user/updatePhone [post]
	UpdatePhone() core.HandlerFunc

	// Follow 关注/取消关注用户
	// @Tags API.user
	// @Router /api/user/follow [post]
	Follow() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	db          mysql.Repo
	cache       redis.Repo
	userService user.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		userService: user.New(db, cache),
	}
}

func (h *handler) i() {}
