package postInfo

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/postInfo"
	"net/http"
)

type getListIdsRequest struct {
	ParentType int32 `json:"parentType" binding:"required"` // 父板块类型
	SubType    int32 `json:"subType" binding:"required"`    // 子版块类型
	SortType   int32 `json:"sortType"`                      // 排序方式 最热 新回复 新发布
}

type getListIdsResponse struct {
	Data []int32 `json:"ids"`
}

// GetListIds 获取前20页帖子ID集合 通过Redis获取，没有则从数据库读取
// @Summary 获取前20页帖子ID集合 通过Redis获取，没有则从数据库读取
// @Description 获取前20页帖子ID集合 通过Redis获取，没有则从数据库读取
// @Tags API.postInfo
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body getListIdsRequest true "请求信息"
// @Success 200 {object} getListIdsResponse
// @Failure 400 {object} code.Failure
// @Router /api/postInfo/ids [get]
func (h *handler) GetListIds() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(getListIdsRequest)
		res := new(getListIdsResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}

		searchData := new(postInfo.GetIdsData)
		searchData.ParentType = req.ParentType
		searchData.SubType = req.SubType
		searchData.SortType = req.SortType

		resListData, err := h.postService.GetListIds(ctx, searchData)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.GetPostIdsListError,
				code.Text(code.GetPostIdsListError)).WithError(err),
			)
			return
		}

		res.Data = resListData
		ctx.Payload(res)
	}
}
