///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package company_preach_info

import (
	"fmt"
	"time"

	"campus-recruiting-api/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *CompanyPreachInfo {
	return new(CompanyPreachInfo)
}

func NewQueryBuilder() *companyPreachInfoQueryBuilder {
	return new(companyPreachInfoQueryBuilder)
}

func (t *CompanyPreachInfo) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type companyPreachInfoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *companyPreachInfoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *companyPreachInfoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&CompanyPreachInfo{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *companyPreachInfoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&CompanyPreachInfo{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *companyPreachInfoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&CompanyPreachInfo{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *companyPreachInfoQueryBuilder) First(db *gorm.DB) (*CompanyPreachInfo, error) {
	ret := &CompanyPreachInfo{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *companyPreachInfoQueryBuilder) QueryOne(db *gorm.DB) (*CompanyPreachInfo, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *companyPreachInfoQueryBuilder) QueryAll(db *gorm.DB) ([]*CompanyPreachInfo, error) {
	var ret []*CompanyPreachInfo
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *companyPreachInfoQueryBuilder) Limit(limit int) *companyPreachInfoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *companyPreachInfoQueryBuilder) Offset(offset int) *companyPreachInfoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereId(p mysql.Predicate, value int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereIdIn(value []int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereIdNotIn(value []int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) OrderById(asc bool) *companyPreachInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereCreatedAtIn(value []time.Time) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) OrderByCreatedAt(asc bool) *companyPreachInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereUpdatedAtIn(value []time.Time) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) OrderByUpdatedAt(asc bool) *companyPreachInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereState(p mysql.Predicate, value int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "state", p),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereStateIn(value []int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "state", "IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereStateNotIn(value []int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "state", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) OrderByState(asc bool) *companyPreachInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "state "+order)
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereCompanyRecruitId(p mysql.Predicate, value int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_recruit_id", p),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereCompanyRecruitIdIn(value []int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_recruit_id", "IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereCompanyRecruitIdNotIn(value []int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_recruit_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) OrderByCompanyRecruitId(asc bool) *companyPreachInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "company_recruit_id "+order)
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WherePreachTime(p mysql.Predicate, value time.Time) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "preach_time", p),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WherePreachTimeIn(value []time.Time) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "preach_time", "IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WherePreachTimeNotIn(value []time.Time) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "preach_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) OrderByPreachTime(asc bool) *companyPreachInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "preach_time "+order)
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereType(p mysql.Predicate, value int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereTypeIn(value []int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereTypeNotIn(value []int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) OrderByType(asc bool) *companyPreachInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereAddress(p mysql.Predicate, value string) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", p),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereAddressIn(value []string) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", "IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereAddressNotIn(value []string) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) OrderByAddress(asc bool) *companyPreachInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "address "+order)
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereSchool(p mysql.Predicate, value string) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "school", p),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereSchoolIn(value []string) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "school", "IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereSchoolNotIn(value []string) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "school", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) OrderBySchool(asc bool) *companyPreachInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "school "+order)
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereIsDeletedIn(value []int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) WhereIsDeletedNotIn(value []int32) *companyPreachInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyPreachInfoQueryBuilder) OrderByIsDeleted(asc bool) *companyPreachInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}
