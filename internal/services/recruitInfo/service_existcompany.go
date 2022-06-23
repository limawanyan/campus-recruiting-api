package recruitInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/company_recruit_info"
)

func (s *service) ExistCompany(ctx core.Context, companyId int32) (exist bool, err error) {
	companyQb := company_recruit_info.NewQueryBuilder()
	count, err := companyQb.WhereId(mysql.EqualPredicate, companyId).
		WhereState(mysql.EqualPredicate, int32(enum.Usable)).
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
