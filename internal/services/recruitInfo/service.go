package recruitInfo

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/company_recruit_info"
	"campus-recruiting-api/internal/repository/redis"
)

type Service interface {
	i()
	// Detail 获取校招详情
	Detail(ctx core.Context, companyId int32) (data *RecruitDetailDate, err error)
	// Follow 关注/取消关注校招公司
	Follow(ctx core.Context, companyId, userId int32) (isFollow bool, err error)
	// ExistCompany 判断公司是否存在
	ExistCompany(ctx core.Context, companyId int32) (bool, error)
	// PageList 分页获取相应校招信息
	PageList(ctx core.Context, searchData *SearchData) (listData []*company_recruit_info.CompanyRecruitInfo, total int64, err error)
	// FollowCompanyIds 用户关注的公司ID集合 key:公司id
	FollowCompanyIds(ctx core.Context, userId int32) (list map[int32]struct{}, err error)
	// ExistFollow 判断用户是否关注
	ExistFollow(ctx core.Context, userId, companyId int32) (bool, error)
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
