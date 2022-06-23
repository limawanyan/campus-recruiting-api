package enum

import "campus-recruiting-api/pkg/errors"

type IndustryType int32

const (
	_                      IndustryType = iota
	Education                           // 教育
	Finance                             // 金融
	ArtificialIntelligence              // 人工智能
	OnlineRetailers                     // 电商
	Game                                // 游戏
	LiveBroadcast                       // 直播/短视频
	InternetIntegration                 // 互联网综合
)

func (t IndustryType) ToStr() (string, error) {
	switch t {
	case Education:
		return "教育", nil
	case Finance:
		return "金融", nil
	case ArtificialIntelligence:
		return "人工智能", nil
	case OnlineRetailers:
		return "电商", nil
	case Game:
		return "游戏", nil
	case LiveBroadcast:
		return "直播|短视频", nil
	case InternetIntegration:
		return "互联网综合", nil
	default:
		return "", errors.New("行业类型错误")
	}
}
