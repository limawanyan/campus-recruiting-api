package enum

import "github.com/pkg/errors"

// FinancingType 融资类型
type FinancingType int32

const (
	_        FinancingType = iota
	Angel                  // 天使轮
	RoundA                 // A轮
	RoundB                 // B轮
	RoundC                 // C轮
	Unfunded               // 未融资
	RoundD                 // D轮及以上
	Listed                 // 已上市
)

// PreachType 宣讲类型
type PreachType int32

const (
	_       PreachType = iota
	OnLine             // 线上
	Offline            // 线下
)

// recruitInfoState 是否可用
type recruitInfoState int32

const (
	_       recruitInfoState = iota
	Usable                   // 可用
	Disable                  // 不可用
)

func (t FinancingType) ToStr() (string, error) {
	switch t {
	case Angel:
		return "天使轮", nil
	case RoundA:
		return "A轮", nil
	case RoundB:
		return "B轮", nil
	case RoundC:
		return "C轮", nil
	case Unfunded:
		return "未融资", nil
	case RoundD:
		return "D轮及以上", nil
	case Listed:
		return "已上市", nil
	default:
		return "", errors.New("融资类型错误")
	}
}

func (t PreachType) ToStr() (string, error) {
	switch t {
	case OnLine:
		return "线上", nil
	case Offline:
		return "线下", nil
	default:
		return "", errors.New("宣讲类型错误")
	}
}
