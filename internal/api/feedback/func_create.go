package feedback

import (
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/feedback"
	"net/http"
)

type createRequest struct {
	Content string `json:"content" binding:"required"`        // 内容
	Tel     string `json:"tel" binding:"required"`            // 电话
	Type    int32  `json:"type" binding:"required,gt=0,lt=4"` // 反馈类型
}

type createResponse struct {
	Id int32 `json:"id"`
}

// Create 用户提交反馈
// @Summary 用户提交反馈
// @Description 用户提交反馈
// @Tags API.feedback
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createRequest true "请求信息"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/feedback/create [post]
func (h *handler) Create() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}

		createDate := new(feedback.CreateFeedbackDate)
		createDate.Tel = req.Tel
		createDate.Content = req.Content
		createDate.Type = req.Type
		createDate.UserId = ctx.SessionUserInfo().UserID

		id, err := h.feedbackService.Create(ctx, createDate)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.FeedbackCreateError,
				code.Text(code.FeedbackCreateError)).WithError(err),
			)
			return
		}

		res.Id = id
		ctx.Payload(res)
	}
}
