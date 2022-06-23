package router

import (
	"campus-recruiting-api/internal/api/dictionary"
	"campus-recruiting-api/internal/api/feedback"
	"campus-recruiting-api/internal/api/postInfo"
	"campus-recruiting-api/internal/api/recruitInfo"
	"campus-recruiting-api/internal/api/salaryInfo"
	"campus-recruiting-api/internal/api/user"
	"campus-recruiting-api/internal/pkg/core"
)

func setApiRouter(r *resource) {

	dictionaryHandler := dictionary.New(r.logger, r.db, r.cache)
	feedbackHandler := feedback.New(r.logger, r.db, r.cache)
	recruitInfoHandler := recruitInfo.New(r.logger, r.db, r.cache)
	salaryInfoHandler := salaryInfo.New(r.logger, r.db, r.cache)
	postInfoHandler := postInfo.New(r.logger, r.db, r.cache)
	userInfoHandler := user.New(r.logger, r.db, r.cache)
	// 无需强制登录,如果登录了则获取用户登录信息
	api := r.mux.Group("/api", core.WrapAuthHandler(r.interceptors.GetLoginInfo))
	{
		// dictionary
		api.GET("/dictionary/subItem/:code", dictionaryHandler.SubItem())

		// feedback
		api.GET("/feedback/list", feedbackHandler.List())

		// recruitInfo
		api.GET("/recruitInfo/:id", recruitInfoHandler.Detail())
		api.GET("/recruitInfo/list", recruitInfoHandler.PagingList())

		// salaryInfo
		api.GET("/salaryInfo/:id", salaryInfoHandler.Detail())
		api.GET("/salaryInfo/list", salaryInfoHandler.PagingList())

		// postInfo
		api.GET("/postInfo/:postId", postInfoHandler.Detail())
		api.GET("/postInfo/comment", postInfoHandler.GetComment())
		api.GET("/postInfo/edit", postInfoHandler.GetEditDetail())
		api.GET("/postInfo/ids", postInfoHandler.GetListIds())
		api.GET("/postInfo/list", postInfoHandler.PagingListByIds())
		api.GET("/postInfo/search", postInfoHandler.PagingListBySearch())

		// userInfo
		api.GET("/user/baseInfo", userInfoHandler.BaseInfo())
		api.POST("/user/login", userInfoHandler.Login())
	}

	// 需要登录验证
	login := r.mux.Group("/api", core.WrapAuthHandler(r.interceptors.CheckLogin))
	{
		// feedback
		login.POST("/feedback/create", feedbackHandler.Create())
		// salaryInfo
		login.POST("/salaryInfo/vote", salaryInfoHandler.Vote())
		login.POST("/salaryInfo/create", salaryInfoHandler.Create())
		// recruitInfo
		login.POST("/recruitInfo/follow", recruitInfoHandler.Follow())
		// postInfo
		login.POST("/postInfo/edit", postInfoHandler.Edit())
		login.POST("/postInfo/delete", postInfoHandler.DeleteByUser())
		login.POST("/postInfo/like", postInfoHandler.LikePost())
		login.POST("/postInfo/comment", postInfoHandler.PutComment())
		login.POST("/postInfo/star", postInfoHandler.StarPost())
		login.GET("/postInfo/user", postInfoHandler.PagingListByUser())
		login.GET("/postInfo/star", postInfoHandler.PagingStarList())
		// userInfo
		login.POST("/user/follow", userInfoHandler.Follow())
		login.POST("/user/updateBase", userInfoHandler.UpdateBaseInfo())
	}
}
