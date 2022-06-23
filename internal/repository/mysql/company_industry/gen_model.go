package company_industry

// CompanyIndustry 公司行业对应表
//go:generate gormgen -structs CompanyIndustry -input .
type CompanyIndustry struct {
	Id       int32 //
	Company  int32 // 公司
	Industry int32 // 行业
}
