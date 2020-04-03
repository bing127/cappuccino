package v1

import (
	"cappuccino/controller/schema"
	"cappuccino/errors"
	"cappuccino/ginplus"
	"cappuccino/middleware"
	"cappuccino/pkg"
	"cappuccino/service"
	"github.com/LyricTian/captcha"
	"github.com/gin-gonic/gin"
)

// Login 用户登录
// @Tags 登录管理
// @Summary 用户登录
// @Param body body schema.LoginParam true "请求参数"
// @Success 200 {object} schema.LoginTokenInfo
// @Failure 400 {object} schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 500 {object} schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/login [post]
func Login(c *gin.Context) {
	var item schema.LoginParam
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}
	if !captcha.VerifyString(item.CaptchaID, item.CaptchaCode) {
		ginplus.ResError(c, errors.New400Response("无效的验证码"))
		return
	}

	user,userErr := service.Verify(c,item.UserName,item.Password)
	if userErr != nil {
		ginplus.ResError(c, errors.New400Response(userErr.Error()))
		return
	}
	// 将用户ID放入上下文
	ginplus.SetUserID(c, user.ID)

	tokenInfo, err := middleware.GenerateToken(*user)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	userInfo := make(map[string]interface{})
	userInfo["token"] = tokenInfo
	userInfo["data"] = user
	pkg.GetRedisInstance().Set("userToken_"+string(user.ID),tokenInfo,0)
	ginplus.ResSuccess(c, userInfo)
}
