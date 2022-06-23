package enum

import "campus-recruiting-api/pkg/errors"

// FeedbackCheckState 反馈审核状态
type FeedbackCheckState int32

const (
	Pending   FeedbackCheckState = iota // 待处理
	Processed                           // 已处理
)

// FeedbackType 反馈类型
type FeedbackType int32

const (
	_            FeedbackType = iota
	Optimization              // 优化建议
	Dysfunction               // 功能障碍
	Complaint                 // 投诉举报
)

func (f FeedbackCheckState) ToStr() (string, error) {
	switch f {
	case Pending:
		return "待处理", nil
	case Processed:
		return "已处理", nil
	default:
		return "", errors.New("审核状态错误")
	}
}

func (f FeedbackType) ToStr() (string, error) {
	switch f {
	case Optimization:
		return "优化建议", nil
	case Dysfunction:
		return "功能障碍", nil
	case Complaint:
		return "投诉举报", nil
	default:
		return "", errors.New("反馈类型错误")
	}
}
