///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package company_industry

import (
	"fmt"

	"campus-recruiting-api/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *CompanyIndustry {
	return new(CompanyIndustry)
}

func NewQueryBuilder() *companyIndustryQueryBuilder {
	return new(companyIndustryQueryBuilder)
}

func (t *CompanyIndustry) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type companyIndustryQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *companyIndustryQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *companyIndustryQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&CompanyIndustry{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *companyIndustryQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&CompanyIndustry{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *companyIndustryQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&CompanyIndustry{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *companyIndustryQueryBuilder) First(db *gorm.DB) (*CompanyIndustry, error) {
	ret := &CompanyIndustry{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *companyIndustryQueryBuilder) QueryOne(db *gorm.DB) (*CompanyIndustry, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *companyIndustryQueryBuilder) QueryAll(db *gorm.DB) ([]*CompanyIndustry, error) {
	var ret []*CompanyIndustry
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *companyIndustryQueryBuilder) Limit(limit int) *companyIndustryQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *companyIndustryQueryBuilder) Offset(offset int) *companyIndustryQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *companyIndustryQueryBuilder) WhereId(p mysql.Predicate, value int32) *companyIndustryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *companyIndustryQueryBuilder) WhereIdIn(value []int32) *companyIndustryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *companyIndustryQueryBuilder) WhereIdNotIn(value []int32) *companyIndustryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyIndustryQueryBuilder) OrderById(asc bool) *companyIndustryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *companyIndustryQueryBuilder) WhereCompany(p mysql.Predicate, value int32) *companyIndustryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company", p),
		value,
	})
	return qb
}

func (qb *companyIndustryQueryBuilder) WhereCompanyIn(value []int32) *companyIndustryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company", "IN"),
		value,
	})
	return qb
}

func (qb *companyIndustryQueryBuilder) WhereCompanyNotIn(value []int32) *companyIndustryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyIndustryQueryBuilder) OrderByCompany(asc bool) *companyIndustryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "company "+order)
	return qb
}

func (qb *companyIndustryQueryBuilder) WhereIndustry(p mysql.Predicate, value int32) *companyIndustryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "industry", p),
		value,
	})
	return qb
}

func (qb *companyIndustryQueryBuilder) WhereIndustryIn(value []int32) *companyIndustryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "industry", "IN"),
		value,
	})
	return qb
}

func (qb *companyIndustryQueryBuilder) WhereIndustryNotIn(value []int32) *companyIndustryQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "industry", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyIndustryQueryBuilder) OrderByIndustry(asc bool) *companyIndustryQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "industry "+order)
	return qb
}
