package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/postInfo"
	"go.uber.org/zap"
	"net/http"
)

type detailRequest struct {
	PostId int32 `json:"postId" uri:"postId" binding:"required"` // 帖子ID
}

type detailResponse struct {
	Data postInfo.PostDetailData `json:"data"`
}

// Detail 根据帖子ID获取帖子详细信息
// @Summary 根据帖子ID获取帖子详细信息
// @Description 根据帖子ID获取帖子详细信息
// @Tags API.postInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/detail/:postId [get]
func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)
		if err := ctx.ShouldBindURI(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		data, err := h.postService.Detail(ctx, req.PostId)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.PostInfoDetailError,
				code.Text(code.PostInfoDetailError)).WithError(err))
			return
		}
		// 增加浏览量
		go func(postId int32) {
			err := h.postService.AddBrows(postId)
			if err != nil {
				h.logger.Error("帖子增加浏览量错误：", zap.Error(err))
			}
		}(req.PostId)
		res.Data = *data

		// 如果为登录状态，获取是否关注 后期可优化 用户登录时一次性获取用户关注列表存储到Redis，相关操作通过取Redis判断
		if ctx.SessionUserInfo().UserID != 0 {
			res.Data.User.IsFollow, err = h.userService.ExistFollow(ctx, ctx.SessionUserInfo().UserID, data.User.ID)
			if err != nil {
				h.logger.Error("获取用户是否关注出错", zap.Error(err))
			}
			// 登录状态判断是否点赞
			res.Data.IsLike, err = h.postService.ExistLike(ctx, ctx.SessionUserInfo().UserID, res.Data.ID, res.Data.ParentType)
			if err != nil {
				h.logger.Error("获取用户是否点赞出错", zap.Error(err))
			}
			// 登录状态判断是否收藏
			res.Data.IsStar, err = h.postService.ExistStar(ctx, ctx.SessionUserInfo().UserID, res.Data.ID, res.Data.ParentType)
			if err != nil {
				h.logger.Error("获取用户是否收藏出错", zap.Error(err))
			}
		}

		ctx.Payload(res)
	}
}
