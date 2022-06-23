package code

import (
	_ "embed"

	"campus-recruiting-api/configs"
)

//go:embed code.go
var ByteCodeFile []byte

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104

	DictionarySubItemError = 20101

	FeedbackCreateError = 20201
	FeedbackListError   = 20202

	RecruitInfoDetailError = 20301
	FollowCompanyError     = 20302
	RecruitInfoListError   = 20303

	SalaryInfoDetailError = 20401
	SalaryVoteError       = 20402
	SalaryInfoCreateError = 20403
	SalaryInfoListError   = 20404

	DeletePostByUserError   = 20501
	PostInfoDetailError     = 20502
	EditPostInfoError       = 20503
	GetCommentListError     = 20504
	GetPostEditDetailError  = 20505
	GetPostIdsListError     = 20506
	LikePostError           = 20507
	GetPostListByIdsError   = 20508
	PostListBySearchError   = 20509
	PagePostListByUserError = 20510
	PagePostStarListError   = 20511
	CreatePostCommentError  = 20512
	StarPostError           = 20513

	FollowUserError         = 20601
	BaseUserInfoError       = 20602
	UpdateUserBaseInfoError = 20603
	WXLoginError            = 20604
	FocusSelfError          = 20605
)

func Text(code int) string {
	lang := configs.Get().Language.Local

	if lang == configs.ZhCN {
		return zhCNText[code]
	}

	if lang == configs.EnUS {
		return enUSText[code]
	}

	return zhCNText[code]
}
