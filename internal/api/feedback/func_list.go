package feedback

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/enum"
	"campus-recruiting-api/internal/pkg/core"
	"net/http"
	"time"
)

type listData struct {
	ID        int32     `json:"id"`        // ID
	Type      string    `json:"type"`      // 反馈类型
	Content   string    `json:"content"`   // 反馈内容
	Tel       string    `json:"tel"`       // 联系电话
	State     string    `json:"state"`     // 反馈状态
	Reply     string    `json:"reply"`     // 回复内容
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 处理时间
}

type listResponse struct {
	List []listData `json:"list"`
}

// List 用户获取反馈记录
// @Summary 用户获取反馈记录
// @Description 用户获取反馈记录
// @Tags API.feedback
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/feedback/list [get]
func (h *handler) List() core.HandlerFunc {
	return func(ctx core.Context) {
		feedbackList, err := h.feedbackService.List(ctx, ctx.SessionUserInfo().UserID)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.FeedbackListError,
				code.Text(code.FeedbackListError)).WithError(err),
			)
			return
		}
		res := new(listResponse)
		res.List = make([]listData, len(feedbackList))
		for i, data := range feedbackList {
			res.List[i].ID = data.Id
			res.List[i].Type, _ = enum.FeedbackType(data.Type).ToStr()
			res.List[i].Reply = data.Reply
			res.List[i].CreatedAt = data.CreatedAt
			res.List[i].UpdatedAt = data.UpdatedAt
			res.List[i].State, _ = enum.FeedbackCheckState(data.State).ToStr()
			res.List[i].Tel = data.Tel
			res.List[i].Content = data.Content
		}
		ctx.Payload(res)
	}
}
