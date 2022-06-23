package dictionary

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/dictionary"
)

type SubItemData struct {
	Label string `json:"label"`
	Value int32  `json:"value"`
}

func (s *service) SubItem(ctx core.Context, code string) (subItemData []SubItemData, err error) {
	dictionaryQb := dictionary.NewQueryBuilder()
	dictionaryQb.WhereTypeCode(mysql.EqualPredicate, code).WhereState(mysql.EqualPredicate, int32(enum.Normal))
	parent, err := dictionaryQb.First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}

	dictionaryQb = dictionary.NewQueryBuilder()
	dictionaryData, err := dictionaryQb.WhereState(mysql.EqualPredicate, int32(enum.Normal)).
		WhereParentId(mysql.EqualPredicate, parent.Id).
		OrderBySort(true).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))

	if err != nil {
		return nil, err
	}

	if len(dictionaryData) <= 0 {
		return
	}

	subItemData = make([]SubItemData, len(dictionaryData))
	for i, v := range dictionaryData {
		subItemData[i].Label = v.Label
		subItemData[i].Value = v.Value
	}

	return
}
