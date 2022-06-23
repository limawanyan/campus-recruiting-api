package user

import (
	"campus-recruiting-api/configs"
	"campus-recruiting-api/internal/code"
	"campus-recruiting-api/internal/pkg/core"
	"campus-recruiting-api/internal/services/user"
	"campus-recruiting-api/pkg/token"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type loginRequest struct {
	Code      string `json:"code" binding:"required"`      // 登录Code
	AvatarUrl string `json:"avatarUrl" binding:"required"` // 用户头像地址
	NickName  string `json:"nickName" binding:"required"`  // 用户昵称
	Gender    int32  `json:"gender" binding:"lt=3"`        // 用户性别
}

type loginResponse struct {
	ID        int32  `json:"id"`        // 用户ID
	Token     string `json:"token"`     // Token令牌
	AvatarUrl string `json:"avatarUrl"` // 用户头像
	NickName  string `json:"nickName"`  // 用户昵称
	Gender    int32  `json:"gender"`    // 用户性别
}

// Login 登录 根据账号密码或微信授权信息验证用户 如果未注册则自动注册一个账号
// @Summary 登录
// @Description 登录
// @Tags API.user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body loginRequest true "请求信息"
// @Success 200 {object} loginResponse
// @Failure 400 {object} code.Failure
// @Router /api/user/login [post]
func (h *handler) Login() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(loginRequest)
		res := new(loginResponse)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err))
			return
		}
		// 获取微信用户OpenId
		url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
			configs.WXAppId,
			configs.WXAppSecret, req.Code)
		wxResp, err := http.DefaultClient.Get(url)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WXLoginError,
				code.Text(code.WXLoginError)).WithError(err))
			return
		}
		var wxMap map[string]string
		err = json.NewDecoder(wxResp.Body).Decode(&wxMap)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WXLoginError,
				code.Text(code.WXLoginError)).WithError(err))
			return
		}
		defer wxResp.Body.Close()
		// 查询该openid用户是否存在数据库中，获取该用户的信息
		userInfo, err := h.userService.GetUserByWxOpenId(ctx, wxMap["openid"])
		if err != nil && err != gorm.ErrRecordNotFound {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WXLoginError,
				code.Text(code.WXLoginError)).WithError(err))
			return
		}
		var tokenClaims token.UserClaims
		// 如果不存在则创建用户并生成token返回
		if err == gorm.ErrRecordNotFound {
			registerData := &user.RegisterUserData{
				NickName:     req.NickName,
				WxOpenId:     wxMap["openid"],
				HeadPortrait: req.AvatarUrl,
				Sex:          req.Gender,
			}
			userId, err := h.userService.Register(ctx, registerData)
			if err != nil {
				ctx.AbortWithError(core.Error(
					http.StatusBadRequest,
					code.WXLoginError,
					code.Text(code.WXLoginError)).WithError(err))
				return
			}
			tokenClaims = token.UserClaims{
				Id:           userId,
				NickName:     req.NickName,
				HeadPortrait: req.AvatarUrl,
				Sex:          req.Gender,
				StandardClaims: jwt.StandardClaims{
					IssuedAt: time.Now().Unix(),
					Issuer:   "SheWaiZhaoPin",
					Subject:  "user token",
				},
			}
			res.ID = userId
			res.AvatarUrl = req.AvatarUrl
			res.Gender = req.Gender
			res.NickName = req.NickName
		} else {
			// 如果存在生成token返回
			tokenClaims = token.UserClaims{
				Id:           userInfo.Id,
				NickName:     userInfo.Nickname,
				HeadPortrait: userInfo.HeadPortrait,
				Sex:          userInfo.Sex,
				StandardClaims: jwt.StandardClaims{
					IssuedAt: time.Now().Unix(),
					Issuer:   "SheWaiZhaoPin",
					Subject:  "user token",
				},
			}
			res.ID = userInfo.Id
			res.AvatarUrl = userInfo.HeadPortrait
			res.Gender = userInfo.Sex
			res.NickName = userInfo.Nickname
		}
		res.Token, err = token.MakeToken(tokenClaims)
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.WXLoginError,
				code.Text(code.WXLoginError)).WithError(err))
			return
		}
		ctx.Payload(res)
	}
}
