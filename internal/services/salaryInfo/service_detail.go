package salaryInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/salary_info"
	"time"
)

type SalaryDetailData struct {
	ID          int32     `json:"id"`
	CompanyName string    `json:"company"`     // 公司名称
	PostName    string    `json:"postName"`    // 岗位名称
	Education   string    `json:"education"`   // 学历
	Reliability int32     `json:"reliability"` // 可信度
	Description string    `json:"description"` // 薪资描述
	CityName    string    `json:"cityName"`    // 城市地点
	CreatedAt   time.Time `json:"createdAt"`   // 发布时间
	Industry    string    `json:"industry"`    // 行业
	Remarks     string    `json:"remarks"`     // 备注信息
}

func (s *service) Detail(ctx core.Context, companyId int32) (data *SalaryDetailData, err error) {
	salaryInfoQb := salary_info.NewQueryBuilder()
	salaryInfoQb.WhereId(mysql.EqualPredicate, companyId).WhereState(mysql.EqualPredicate, int32(enum.Usable))
	salaryInfoData, err := salaryInfoQb.First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	data = &SalaryDetailData{
		ID:          salaryInfoData.Id,
		CompanyName: salaryInfoData.CompanyName,
		PostName:    salaryInfoData.PostName,
		Reliability: salaryInfoData.Reliability,
		Description: salaryInfoData.Description,
		CityName:    salaryInfoData.CityName,
		CreatedAt:   salaryInfoData.CreatedAt,
		Remarks:     salaryInfoData.Remarks,
	}
	data.Education, _ = enum.EducationType(salaryInfoData.Education).ToStr()
	data.Industry, _ = enum.IndustryType(salaryInfoData.Industry).ToStr()
	return
}
