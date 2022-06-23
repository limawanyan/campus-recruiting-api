package enum

import "campus-recruiting-api/pkg/errors"

type EducationType int32

const (
	_                             EducationType = iota
	JuniorCollege                               // 大专
	UndergraduateOther                          // 本科其他
	UndergraduateReturnees                      // 本科海归
	Undergraduate211                            // 本科211
	Undergraduate985                            // 本科985
	UndergraduateDoubleFirstClass               // 本科双一流
	MasterOther                                 // 硕士其他
	MasterReturnees                             // 硕士海归
	Master211                                   // 硕士211
	MasterDoubleFirstClass                      // 硕士双一流
	doctorOther                                 // 博士其他
	doctorReturnees                             // 博士海归
	doctor211                                   // 博士211
	doctor985                                   // 博士985
	doctorDoubleFirstClass                      // 博士双一流
	Other                                       // 其他
)

func (t EducationType) ToStr() (string, error) {
	switch t {
	case JuniorCollege:
		return "大专", nil
	case UndergraduateOther:
		return "本科其他", nil
	case UndergraduateReturnees:
		return "本科海归", nil
	case Undergraduate211:
		return "本科211", nil
	case Undergraduate985:
		return "本科985", nil
	case UndergraduateDoubleFirstClass:
		return "本科双一流", nil
	case MasterOther:
		return "硕士其他", nil
	case MasterReturnees:
		return "硕士海归", nil
	case Master211:
		return "硕士211", nil
	case MasterDoubleFirstClass:
		return "硕士双一流", nil
	case doctorOther:
		return "博士其他", nil
	case doctorReturnees:
		return "博士海归", nil
	case doctor211:
		return "博士211", nil
	case doctor985:
		return "博士985", nil
	case doctorDoubleFirstClass:
		return "博士双一流", nil
	default:
		return "", errors.New("学历类型错误")
	}
}
