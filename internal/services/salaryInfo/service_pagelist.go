package salaryInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/salary_info"
)

type SearchData struct {
	Page      int    `json:"page"`      // 第几页
	PageSize  int    `json:"pageSize"`  // 每页显示条数
	Industry  int32  `json:"industry"`  // 所属行业
	Education int32  `json:"education"` // 公司属性
	KeyWord   string `json:"keyWord"`   // 模糊搜索关键字
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (listData []*salary_info.SalaryInfo, total int64, err error) {
	page := searchData.Page
	if page <= 0 {
		page = 1
	}
	pageSize := searchData.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	salaryInfoQb := salary_info.NewQueryBuilder()
	salaryInfoQb.WhereReliability(mysql.GreaterThanPredicate, enum.Critical).
		WhereState(mysql.EqualPredicate, int32(enum.Usable))

	if searchData.Industry != 0 {
		salaryInfoQb.WhereIndustry(mysql.EqualPredicate, searchData.Industry)
	}
	if searchData.Education != 0 {
		salaryInfoQb.WhereEducation(mysql.EqualPredicate, searchData.Education)
	}
	if searchData.KeyWord != "" {
		salaryInfoQb.WhereFuzzyQuery(mysql.LikePredicate, "%"+searchData.KeyWord+"%")
	}

	total, err = salaryInfoQb.Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, 0, err
	}

	listData, err = salaryInfoQb.
		Limit(pageSize).
		Offset(offset).
		OrderByCreatedAt(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, 0, err
	}

	return
}
