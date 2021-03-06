///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package company_recruit_info

import (
	"fmt"
	"time"

	"campus-recruiting-api/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *CompanyRecruitInfo {
	return new(CompanyRecruitInfo)
}

func NewQueryBuilder() *companyRecruitInfoQueryBuilder {
	return new(companyRecruitInfoQueryBuilder)
}

func (t *CompanyRecruitInfo) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type companyRecruitInfoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *companyRecruitInfoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *companyRecruitInfoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&CompanyRecruitInfo{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *companyRecruitInfoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&CompanyRecruitInfo{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *companyRecruitInfoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&CompanyRecruitInfo{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		return 0, nil
	}
	return c, res.Error
}

func (qb *companyRecruitInfoQueryBuilder) First(db *gorm.DB) (*CompanyRecruitInfo, error) {
	ret := &CompanyRecruitInfo{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *companyRecruitInfoQueryBuilder) QueryOne(db *gorm.DB) (*CompanyRecruitInfo, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *companyRecruitInfoQueryBuilder) QueryAll(db *gorm.DB) ([]*CompanyRecruitInfo, error) {
	var ret []*CompanyRecruitInfo
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *companyRecruitInfoQueryBuilder) Limit(limit int) *companyRecruitInfoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) Offset(offset int) *companyRecruitInfoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereId(p mysql.Predicate, value int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereIdIn(value []int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereIdNotIn(value []int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderById(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCreatedAtIn(value []time.Time) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByCreatedAt(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereUpdatedAtIn(value []time.Time) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByUpdatedAt(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereDescription(p mysql.Predicate, value string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereDescriptionIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereDescriptionNotIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "description", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByDescription(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "description "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereState(p mysql.Predicate, value int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "state", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereStateIn(value []int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "state", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereStateNotIn(value []int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "state", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByState(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "state "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCompanyName(p mysql.Predicate, value string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_name", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCompanyNameIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_name", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCompanyNameNotIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByCompanyName(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "company_name "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCompanyAttribute(p mysql.Predicate, value int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_attribute", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCompanyAttributeIn(value []int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_attribute", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCompanyAttributeNotIn(value []int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_attribute", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByCompanyAttribute(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "company_attribute "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereRecommendTime(p mysql.Predicate, value string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_time", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereRecommendTimeIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_time", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereRecommendTimeNotIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "recommend_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByRecommendTime(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "recommend_time "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereApplyOnlineTime(p mysql.Predicate, value string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "apply_online_time", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereApplyOnlineTimeIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "apply_online_time", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereApplyOnlineTimeNotIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "apply_online_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByApplyOnlineTime(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "apply_online_time "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereWrittenExamTime(p mysql.Predicate, value string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "written_exam_time", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereWrittenExamTimeIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "written_exam_time", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereWrittenExamTimeNotIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "written_exam_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByWrittenExamTime(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "written_exam_time "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereInterviewTime(p mysql.Predicate, value string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "interview_time", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereInterviewTimeIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "interview_time", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereInterviewTimeNotIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "interview_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByInterviewTime(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "interview_time "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereOfferTime(p mysql.Predicate, value string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "offer_time", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereOfferTimeIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "offer_time", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereOfferTimeNotIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "offer_time", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByOfferTime(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "offer_time "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereIndustrys(p mysql.Predicate, value string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "industrys", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereIndustrysIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "industrys", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereIndustrysNotIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "industrys", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByIndustrys(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "industrys "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCompanyLogo(p mysql.Predicate, value string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_logo", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCompanyLogoIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_logo", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereCompanyLogoNotIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "company_logo", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByCompanyLogo(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "company_logo "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereAddress(p mysql.Predicate, value string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereAddressIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereAddressNotIn(value []string) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByAddress(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "address "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereFinancingType(p mysql.Predicate, value int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "financing_type", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereFinancingTypeIn(value []int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "financing_type", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereFinancingTypeNotIn(value []int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "financing_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByFinancingType(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "financing_type "+order)
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereIsDeletedIn(value []int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) WhereIsDeletedNotIn(value []int32) *companyRecruitInfoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *companyRecruitInfoQueryBuilder) OrderByIsDeleted(asc bool) *companyRecruitInfoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}
