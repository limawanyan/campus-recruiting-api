package configs

import "time"

const (
	// ProjectName 项目名称
	ProjectName = "campus-recruiting-api"

	// ProjectDomain 项目域名
	ProjectDomain = "http://127.0.0.1"

	// ProjectPort 项目端口
	ProjectPort = ":8081"

	// ProjectAccessLogFile 项目访问日志存放位置
	ProjectAccessLogFile = "./logs/" + ProjectName + "-access.log"

	HeaderLoginToken = "Token"

	// HeaderSignToken 签名验证 Authorization，Header 中传递的参数
	HeaderSignToken = "Authorization"

	// HeaderSignTokenDate 签名验证 Date，Header 中传递的参数
	HeaderSignTokenDate = "Authorization-Date"

	// MaxRequestsPerSecond 每秒最大请求量
	MaxRequestsPerSecond = 10000

	// ZhCN 简体中文 - 中国
	ZhCN = "zh-cn"

	// EnUS 英文 - 美国
	EnUS = "en-us"

	// LoginSessionTTL 登录有效期为 24 小时
	LoginSessionTTL = time.Hour * 24

	// NotExpired Redis Key不过期
	NotExpired = 0

	// PageSize 首页帖子默认一页多少条记录
	PageSize = 10

	// IdsMaxPage 最大获取页数
	IdsMaxPage = 20

	// HotValue 一个点赞 += 热门权重432
	HotValue = 432

	// WXAppId 微信小程序应用ID WXAppSecret 小程序密钥
	WXAppId     = "wxacc09df80ea695b5"
	WXAppSecret = "04a0fa6ca1e67a4c1e5d262254f6e489"

	// JwtSigningKey JWT密钥
	JwtSigningKey = "dsafeafaafadsf"
)
