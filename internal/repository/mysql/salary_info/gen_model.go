package salary_info

import "time"

// SalaryInfo 薪资爆料信息
//go:generate gormgen -structs SalaryInfo -input .
type SalaryInfo struct {
	Id          int32     // 数据唯一标识
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	Description string    // 薪资描述（薪资详情）
	State       int32     // 状态，是否上下架
	CompanyName string    // 公司名称
	PostName    string    // 岗位名称
	CityName    string    // 城市名称
	Remarks     string    // 备注（offer相关信息）
	Education   int32     // 学历
	Industry    int32     // 行业
	Reliability int32     // 可信度
	FromUserId  int32     // 爆料人
	FuzzyQuery  string    // 模糊查询（公司+职位+地点）
}
