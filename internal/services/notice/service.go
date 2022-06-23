package notice

import (
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/redis"
)

type Service interface {
	i()
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
