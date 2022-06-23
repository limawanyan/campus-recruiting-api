package postInfo

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/redis"
	"campus-recruiting-api/internal/services/postInfo"
	"campus-recruiting-api/internal/services/user"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// GetListIds 获取前20页帖子ID集合 通过Redis获取，没有则从数据库读取
	// @Tags API.postInfo
	// @Router /api/postInfo/ids [get]
	GetListIds() core.HandlerFunc

	// PagingListByIds 根据ID集合获取分页数据列表
	// @Tags API.postInfo
	// @Router /api/postInfo/list [get]
	PagingListByIds() core.HandlerFunc

	// PagingListBySearch 模糊搜索分页获取面经/面题，从数据库获取，不限制最多为前20页
	// @Tags API.postInfo
	// @Router /api/postInfo/search [get]
	PagingListBySearch() core.HandlerFunc

	// LikePost 点赞帖子
	// @Tags API.postInfo
	// @Router /api/postInfo/like [post]
	LikePost() core.HandlerFunc

	// StarPost 收藏帖子
	// @Tags API.postInfo
	// @Router /api/postInfo/star [post]
	StarPost() core.HandlerFunc

	// Detail 根据帖子ID获取帖子详细信息
	// @Tags API.postInfo
	// @Router /api/postInfo/:postId [get]
	Detail() core.HandlerFunc

	// GetComment 获取帖子评论
	// @Tags API.postInfo
	// @Router /api/postInfo/comment [get]
	GetComment() core.HandlerFunc

	// PutComment 发表评论
	// @Tags API.postInfo
	// @Router /api/postInfo/comment [post]
	PutComment() core.HandlerFunc

	// PagingListByUser 分页获取用户发布的面经/面题,按时间倒序
	// @Tags API.postInfo
	// @Router /api/postInfo/user [get]
	PagingListByUser() core.HandlerFunc

	// DeleteByUser 用户删除自己发布的面经/面题
	// @Tags API.postInfo
	// @Router /api/postInfo/delete [post]
	DeleteByUser() core.HandlerFunc

	// GetEditDetail 根据帖子ID获取帖子详细信息（标题、类别、内容）
	// @Tags API.postInfo
	// @Router /api/postInfo/edit [get]
	GetEditDetail() core.HandlerFunc

	// Edit 用户发布或重新编辑帖子
	// @Tags API.postInfo
	// @Router /api/postInfo/edit [post]
	Edit() core.HandlerFunc

	// PagingStarList 分页获取用户收藏的面经/面题
	// @Tags API.postInfo
	// @Router /api/postInfo/star [get]
	PagingStarList() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	db          mysql.Repo
	cache       redis.Repo
	postService postInfo.Service
	userService user.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:      logger,
		cache:       cache,
		postService: postInfo.New(db, cache),
		userService: user.New(db, cache),
	}
}

func (h *handler) i() {}
