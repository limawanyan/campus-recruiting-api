package feedback

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/user_feedback"
	"campus-recruiting-api/internal/repository/redis"
)

type Service interface {
	i()
	// Create 创建反馈
	Create(ctx core.Context, createDate *CreateFeedbackDate) (id int32, err error)
	// List 获取用户反馈记录列表
	List(ctx core.Context, userId int32) (listData []*user_feedback.UserFeedback, err error)
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
