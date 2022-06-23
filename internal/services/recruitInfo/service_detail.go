package recruitInfo

import (
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/mysql/company_preach_info"
	"campus-recruiting-api/internal/repository/mysql/company_recruit_info"
)

type RecruitDetailDate struct {
	ID              int32        `json:"id"`
	Name            string       `json:"name"`            // 公司名称
	RecommendTime   string       `json:"recommendTime"`   // 内推时间
	ApplyOnlineTime string       `json:"applyOnlineTime"` // 网申时间
	WrittenExamTime string       `json:"writtenExamTime"` // 笔试时间
	InterviewTime   string       `json:"interviewTime"`   // 面试时间
	OfferTime       string       `json:"offerTime"`       // Offer发放时间
	Industry        string       `json:"industry"`        // 所属行业
	Logo            string       `json:"logo"`            // Logo
	Address         string       `json:"address"`         // 地址
	FinancingType   string       `json:"financingType"`   // 融资类型
	PreachList      []PreachData `json:"preachList"`      // 宣讲信息列表
	IsFollow        bool         `json:"isFollow"`        // 是否关注
}

type PreachData struct {
	Type    string `json:"type"`    // 宣讲类型
	Address string `json:"address"` // 宣讲地点
	School  string `json:"school"`  // 宣讲学校名称
}

func (s *service) Detail(ctx core.Context, companyId int32) (data *RecruitDetailDate, err error) {
	recruitInfoQb := company_recruit_info.NewQueryBuilder()
	recruitInfoQb.WhereId(mysql.EqualPredicate, companyId).WhereState(mysql.EqualPredicate, int32(enum.Usable))
	recruitInfoData, err := recruitInfoQb.First(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	preachList, err := company_preach_info.NewQueryBuilder().
		WhereCompanyRecruitId(mysql.EqualPredicate, companyId).
		WhereState(mysql.EqualPredicate, int32(enum.Usable)).
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	data = &RecruitDetailDate{}
	data.ID = recruitInfoData.Id
	data.Address = recruitInfoData.Address
	data.Name = recruitInfoData.CompanyName
	data.RecommendTime = recruitInfoData.RecommendTime
	data.ApplyOnlineTime = recruitInfoData.ApplyOnlineTime
	data.WrittenExamTime = recruitInfoData.WrittenExamTime
	data.InterviewTime = recruitInfoData.InterviewTime
	data.OfferTime = recruitInfoData.OfferTime
	data.Industry = recruitInfoData.Industrys
	data.Logo = recruitInfoData.CompanyLogo
	data.FinancingType, _ = enum.FinancingType(recruitInfoData.FinancingType).ToStr()
	data.PreachList = make([]PreachData, len(preachList))
	for i, v := range preachList {
		data.PreachList[i].Type, _ = enum.PreachType(v.Type).ToStr()
		data.PreachList[i].Address = v.Address
		data.PreachList[i].School = v.School
	}
	return
}
