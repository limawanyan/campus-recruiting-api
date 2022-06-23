package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/postInfo"
	"net/http"
	"time"
)

type editRequest struct {
	ID         int32  `json:"id"`
	Title      string `json:"title" binding:"required"`      // 标题
	Content    string `json:"content" binding:"required"`    // 内容
	ParentType int32  `json:"parentType" binding:"required"` // 父级板块类型
	SubType    int32  `json:"subType" binding:"required"`    // 子版块类型
}

type editResponse struct {
	IsCreate bool `json:"isCreate"` // true 创建 false 编辑
}

// Edit 用户发布或重新编辑帖子
// @Summary 用户发布或重新编辑帖子
// @Description 用户发布或重新编辑帖子
// @Tags API.postInfo
// @Accept application/x-www-form-url	encoded
// @Produce json
// @Param Request body editRequest true "请求信息"
// @Success 200 {object} editResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/edit [post]
func (h *handler) Edit() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(editRequest)
		res := new(editResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		// 待完善 父子级板块类型校验

		data := new(postInfo.EditData)
		data.ID = req.ID
		data.Title = req.Title
		data.Content = req.Content
		data.ParentType = req.ParentType
		data.SubType = req.SubType
		data.FromUserId = ctx.SessionUserInfo().UserID
		if req.ID == 0 {
			res.IsCreate = true
			data.SortWeight = time.Now().Unix() // 热门排序权重 默认当前时间戳
		}
		err := h.postService.Edit(ctx, data)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.EditPostInfoError,
				code.Text(code.EditPostInfoError)).WithError(err),
			)
			return
		}

		ctx.Payload(res)
	}
}
