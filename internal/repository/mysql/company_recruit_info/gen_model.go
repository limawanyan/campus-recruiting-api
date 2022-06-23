package company_recruit_info

import "time"

// CompanyRecruitInfo 公司校招信息
//go:generate gormgen -structs CompanyRecruitInfo -input .
type CompanyRecruitInfo struct {
	Id               int32     // 数据唯一标识
	CreatedAt        time.Time `gorm:"time"` // 创建时间
	UpdatedAt        time.Time `gorm:"time"` // 更新时间
	Description      string    // 描述信息
	State            int32     // 状态，是否可用
	CompanyName      string    // 公司名称
	CompanyAttribute int32     // 公司属性
	RecommendTime    string    // 内推时间
	ApplyOnlineTime  string    // 网申时间
	WrittenExamTime  string    // 笔试时间（可能存在多个时间）
	InterviewTime    string    // 面试时间
	OfferTime        string    // offer发放时间
	Industrys        string    // 所属行业名称(多个行业顿号拼接)
	CompanyLogo      string    // 公司logo
	Address          string    // 地址(多个地址顿号拼接)
	FinancingType    int32     // 融资类型
	IsDeleted        int32     // 是否逻辑删除（0 否 1是）
}
