package user

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user"
	"campus-recruiting-api/internal/repository/redis"
)

type Service interface {
	i()

	// BaseInfo 获取用户基础信息
	BaseInfo(ctx core.Context, userId int32) (data *BaseInfoData, err error)
	// ExistUser 用户是否存在
	ExistUser(ctx core.Context, userId int32) (exist bool, err error)
	// Follow 关注/取消关注用户
	Follow(ctx core.Context, userId, fromUserId int32) (isFollow bool, err error)
	// UpdateBaseInfo 更新用户基础信息
	UpdateBaseInfo(ctx core.Context, updateData *UpdateBaseInfoData) error
	// ExistFollow 判断用户A是否关注用户B
	ExistFollow(ctx core.Context, fromUserId, toUserId int32) (exist bool, err error)
	// FollowNum 用户关注数量
	FollowNum(ctx core.Context, userId int32) (count int64, err error)
	// FansNum 用户粉丝数量
	FansNum(ctx core.Context, userId int32) (count int64, err error)
	// GetUserByWxOpenId 获取用户信息
	GetUserByWxOpenId(ctx core.Context, openId string) (data *user.User, err error)
	// Register 注册
	Register(ctx core.Context, data *RegisterUserData) (id int32, err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {

}
