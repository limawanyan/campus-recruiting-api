package dictionary

import "time"

// Dictionary 数据字典;管理网站类型数据 行业类型 公司属性 讨论帖类型 面经类型
//go:generate gormgen -structs Dictionary -input .
type Dictionary struct {
	Id          int32     // 记录ID
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	Name        string    // 字典名称
	TypeCode    string    // 字典类型标识
	Label       string    // 子项展示值
	Value       int32     // 子项真实值
	State       int32     // 状态，是否可用
	Description string    // 描述信息
	Sort        int32     // 排序
	ParentId    int32     // 父级ID
	Readonly    int32     // 是否只读，系统初始化固定值，不支持修改（默认0不可修改,1可修改）
	IsDeleted   int32     // 是否逻辑删除（0 否 1是）
}
