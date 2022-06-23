package feedback

import (
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/repository/mysql/user_feedback"
)

type CreateFeedbackDate struct {
	UserId  int32  // 反馈用户
	Content string // 内容
	Tel     string // 电话
	Type    int32  // 反馈类型
}

func (s *service) Create(ctx core.Context, createDate *CreateFeedbackDate) (id int32, err error) {
	model := user_feedback.NewModel()
	model.UserId = createDate.UserId
	model.Content = createDate.Content
	model.Tel = createDate.Tel
	model.Type = createDate.Type

	id, err = model.Create(s.db.GetDbW().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, nil
	}
	return
}
