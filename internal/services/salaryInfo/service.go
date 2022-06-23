package salaryInfo

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/salary_info"
	"campus-recruiting-api/internal/repository/redis"
)

type Service interface {
	i()
	// PageList 分页获取薪资爆料信息
	PageList(ctx core.Context, searchData *SearchData) (listData []*salary_info.SalaryInfo, total int64, err error)
	// Detail 薪资爆料详情
	Detail(ctx core.Context, companyId int32) (data *SalaryDetailData, err error)
	// Vote 可信度投票
	Vote(ctx core.Context, salaryId, userId int32, isSupport bool) (isRepeat bool, err error)
	// Create 爆料薪资
	Create(ctx core.Context, data *CreateSalaryInfoData) (id int32, err error)
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
