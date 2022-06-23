package recruitInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/company_industry"
	"campus-recruiting-api/internal/repository/mysql/company_recruit_info"
	"campus-recruiting-api/internal/repository/mysql/user_follow"
)

type SearchData struct {
	Page      int   `json:"page"`         // 第几页
	PageSize  int   `json:"pageSize"`     // 每页显示条数
	Industry  int32 `json:"industryType"` // 所属行业
	Attribute int32 `json:"attribute"`    // 公司属性
	IsFollow  bool  `json:"isFollow"`     // 用户是否关注
	UserId    int32 `json:"userId"`       // 当前用户ID
}

func (s *service) PageList(ctx core.Context, searchData *SearchData) (listData []*company_recruit_info.CompanyRecruitInfo, total int64, err error) {
	// 如果选择了所属行业，则根据行业类型去 公司行业类型对照表中查询 获取到该行业的所有公司ID
	// 如果选择了用户关注 则通过用户ID去关注表中查询用户关注的所有公司
	page := searchData.Page
	if page <= 0 {
		page = 1
	}
	pageSize := searchData.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	qb := company_recruit_info.NewQueryBuilder()
	qb.WhereState(mysql.EqualPredicate, int32(enum.Usable))

	if searchData.IsFollow {
		follows, err := user_follow.NewQueryBuilder().
			WhereType(mysql.EqualPredicate, int32(enum.Company)).
			WhereFromUserId(mysql.EqualPredicate, searchData.UserId).
			QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
		if err != nil {
			return nil, 0, err
		}
		followIds := make([]int32, len(follows))
		for i, v := range follows {
			followIds[i] = v.TopicId
		}
		qb.WhereIdIn(followIds)
	}

	if searchData.Attribute != 0 {
		qb.WhereCompanyAttribute(mysql.EqualPredicate, searchData.Attribute)
	}

	if searchData.Industry != 0 {
		companyByIndustry, err := company_industry.NewQueryBuilder().
			WhereIndustry(mysql.EqualPredicate, searchData.Industry).
			QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
		if err != nil {
			return nil, 0, err
		}
		companyByIndustryIds := make([]int32, len(companyByIndustry))
		for i, v := range companyByIndustry {
			companyByIndustryIds[i] = v.Company
		}
		qb.WhereIdIn(companyByIndustryIds)
	}

	total, err = qb.Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, 0, err
	}

	listData, err = qb.
		Limit(pageSize).
		Offset(offset).
		OrderById(false).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, 0, err
	}

	return
}
