package dictionary

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/dictionary"
)

func (s *service) IsSubItem(ctx core.Context, code string, value int32) (bool, error) {
	dictionaryQb := dictionary.NewQueryBuilder()
	dictionaryQb.WhereTypeCode(mysql.EqualPredicate, code).WhereState(mysql.EqualPredicate, int32(enum.Normal))
	parent, err := dictionaryQb.First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return false, err
	}

	dictionaryQb = dictionary.NewQueryBuilder()
	count, err := dictionaryQb.WhereState(mysql.EqualPredicate, int32(enum.Normal)).
		WhereParentId(mysql.EqualPredicate, parent.Id).
		WhereValue(mysql.EqualPredicate, value).
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))

	if err != nil {
		return false, err
	}

	if count <= 0 {
		return false, nil
	}

	return true, nil

}
