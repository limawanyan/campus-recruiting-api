package salaryInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/like_info"
	"campus-recruiting-api/internal/repository/mysql/salary_info"
)

func (s *service) Vote(ctx core.Context, salaryId, userId int32, isSupport bool) (isRepeat bool, err error) {
	salaryInfoQb := salary_info.NewQueryBuilder()
	// 判断爆料记录是否存在
	_, err = salaryInfoQb.
		WhereState(mysql.EqualPredicate, int32(enum.Usable)).
		WhereId(mysql.EqualPredicate, salaryId).
		WhereReliability(mysql.GreaterThanPredicate, enum.Critical). // 可信度要大于0
		First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return false, err
	}

	likeInfoQb := like_info.NewQueryBuilder()
	count, err := likeInfoQb.
		WhereType(mysql.EqualPredicate, int32(enum.SalaryDisclosure)).
		WhereTopicId(mysql.EqualPredicate, salaryId).
		WhereFromUid(mysql.EqualPredicate, userId).
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return false, err
	}
	// 已存在投票记录,不可重复投票
	if count > 0 {
		isRepeat = true
		return
	}

	// 投票记录
	likeData := like_info.NewModel()
	likeData.TopicId = salaryId
	likeData.Type = int32(enum.SalaryDisclosure)
	likeData.FromUid = userId
	likeData.State = int32(enum.Untrusted)
	if isSupport {
		likeData.State = int32(enum.Credible)
	}
	_, err = likeData.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return
	}
	updateSql := "UPDATE salary_info SET reliability = reliability-1 WHERE id = ?"
	if isSupport {
		updateSql = "UPDATE salary_info SET reliability = reliability+1 WHERE id = ?"
	}
	// 更新可信度
	err = s.db.GetDbW().Exec(updateSql, salaryId).Error
	if err != nil {
		return
	}

	return
}
