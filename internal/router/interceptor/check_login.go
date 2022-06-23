package interceptor

import (
	"campus-recruiting-api/configs"
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/proposal"
	"campus-recruiting-api/pkg/token"
	"github.com/pkg/errors"
	"net/http"
)

func (i *interceptor) CheckLogin(ctx core.Context) (info proposal.SessionUserInfo, error core.BusinessError) {
	tokenStr := ctx.GetHeader(configs.HeaderSignToken)
	if tokenStr == "" {
		error = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(errors.New("Header中缺少Token参数"))
		return
	}
	user, err := token.ParseToken(tokenStr)
	if err != nil {
		error = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(errors.New("Header中Token校验错误"))
		return
	}
	userInfo := proposal.SessionUserInfo{
		UserID:       user.Id,
		UserName:     user.NickName,
		HeadPortrait: user.HeadPortrait,
		Sex:          user.Sex,
	}
	return userInfo, nil
}
