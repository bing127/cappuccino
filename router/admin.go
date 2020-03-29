package router

import (
	"cappuccino/config"
	v1 "cappuccino/controller/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initAdminRouter(r *gin.Engine) {
	// 后台管理系统
	r.GET(config.Admin.App.ApiPrefix+"/admin", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title":config.Admin.App.Name,
		})
	})

	api := r.Group("/api")

	apiV1 := api.Group("/v1")
	{
		// 系统版本
		apiV1.GET("", v1.V1)
		// 系统登陆相关
		apiV1Login := apiV1.Group("/login")
		{
			// 获取验证码ID  /api/v1/login/captchaId
			apiV1Login.GET("captchaId",v1.GetCaptchaId)
			// 获取验证码图片 /api/v1/login/captcha
			apiV1Login.GET("captcha",v1.ResCaptcha)
			// 登陆  /api/v1/login
			apiV1Login.POST("",v1.Login)
		}
	}


	//r.POST(config.Admin.App.ApiPrefix+"/api/admin/getAccessToken", v1.GetAccessToken)


}