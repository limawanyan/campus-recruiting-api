package salaryInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql/salary_info"
)

type CreateSalaryInfoData struct {
	CompanyName string // 公司名称
	PostName    string // 岗位名称
	CityName    string // 城市名称
	Remarks     string // 备注
	Description string // 薪资描述
	Education   int32  // 学历
	Industry    int32  // 行业
	FromUserId  int32  // 爆料人
	FuzzyQuery  string // 模糊搜索关键词
}

func (s *service) Create(ctx core.Context, data *CreateSalaryInfoData) (id int32, err error) {
	model := salary_info.NewModel()
	model.Description = data.Description
	model.CompanyName = data.CompanyName
	model.PostName = data.PostName
	model.CityName = data.CityName
	model.Remarks = data.Remarks
	model.Education = data.Education
	model.Industry = data.Industry
	model.FromUserId = data.FromUserId
	model.FuzzyQuery = data.FuzzyQuery
	model.State = int32(enum.Usable)
	model.Reliability = 3
	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, nil
	}
	return
}
