package postInfo

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql"
	"campus-recruiting-api/internal/repository/redis"
)

type Service interface {
	i()

	// DeleteByUser 用户删除自己发布帖子
	DeleteByUser(ctx core.Context, postId, userId int32) error
	// Detail 获取帖子详情
	Detail(ctx core.Context, postId int32) (data *PostDetailData, err error)
	// AddBrows 增加帖子浏览量
	AddBrows(postId int32) (err error)
	// Edit 发布/编辑帖子
	Edit(ctx core.Context, data *EditData) error
	// PageCommentList 分页获取评论
	PageCommentList(ctx core.Context, pageData *GetPageCommentData) (data []*PageCommentList, total int64, err error)
	// EditDetail 获取编辑内容
	EditDetail(ctx core.Context, postId int32) (data *EditDetailData, err error)
	// GetListIds 获取ID集合
	GetListIds(ctx core.Context, searchData *GetIdsData) (data []int32, err error)
	// ExistPost 帖子是否存在
	ExistPost(ctx core.Context, postId int32) (exist bool, err error)
	// Like 点赞/取消赞
	Like(ctx core.Context, postId, likeType, userId int32) (isLike bool, err error)
	// ListByIds 根据ID集合获取帖子信息
	ListByIds(ctx core.Context, Ids []int32) (data []PostListData, err error)
	// PageListBySearch 根据关键字模糊搜索帖子
	PageListBySearch(ctx core.Context, searchData *SearchData) (data []PostListData, total int64, err error)
	// PageListByUser 获取当前用户发布的面经/面题
	PageListByUser(ctx core.Context, searchData *SearchByUserData) (data []PostListData, total int64, err error)
	// PageStarList 分页获取用户收藏的面经/面题
	PageStarList(ctx core.Context, searchData *SearchByUserData) (data []PostListData, total int64, err error)
	// CreateComment 创建评论
	CreateComment(ctx core.Context, createData *CreateCommentData) (id int32, err error)
	// Star 收藏/取消收藏帖子
	Star(ctx core.Context, postId, likeType, userId int32) (isStar bool, err error)
	// ExistLike 用户是否点赞
	ExistLike(ctx core.Context, fromUserId, postId, postType int32) (exist bool, err error)
	// ExistStar 用户是否收藏
	ExistStar(ctx core.Context, fromUserId, postId, postType int32) (exist bool, err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {

}
