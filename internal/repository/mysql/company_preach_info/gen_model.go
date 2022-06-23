package company_preach_info

import "time"

// CompanyPreachInfo 公司宣讲信息
//go:generate gormgen -structs CompanyPreachInfo -input .
type CompanyPreachInfo struct {
	Id               int32     // 数据唯一标识
	CreatedAt        time.Time `gorm:"time"` // 创建时间
	UpdatedAt        time.Time `gorm:"time"` // 更新时间
	State            int32     // 状态，是否可用
	CompanyRecruitId int32     // 所属校招公司
	PreachTime       time.Time `gorm:"time"` // 宣讲时间
	Type             int32     // 宣讲类型(在线/线下)
	Address          string    // 宣讲详细地址
	School           string    // 宣讲学校
	IsDeleted        int32     // 是否逻辑删除（0 否 1是）
}
