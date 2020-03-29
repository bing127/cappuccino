package v1

import (
	"cappuccino/config"
	"cappuccino/errors"
	"cappuccino/ginplus"
	"github.com/LyricTian/captcha"
	"github.com/gin-gonic/gin"
)


// GetCaptcha 获取验证码信息
// @Tags 登录管理
// @Summary 获取验证码信息
// @Success 200 {string} string
// @Router /api/v1/login/captchaId [get]
func GetCaptchaId(c *gin.Context) {
	captchaID := captcha.NewLen(config.Admin.Captcha.Length)
	result := make(map[string]interface{})
	result["captchaID"] = captchaID
	ginplus.ResSuccess(c, result)
}

// ResCaptcha 响应图形验证码
// @Tags 登录管理
// @Summary 响应图形验证码
// @Param id query string true "验证码ID"
// @Produce image/png
// @Success 200 "图形验证码"
// @Failure 400 {object} schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 500 {object} schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/login/captcha [get]
func ResCaptcha(c *gin.Context) {
	captchaID := c.Query("id")
	if captchaID == "" {
		ginplus.ResError(c, errors.New400Response("请提供验证码ID"))
		return
	}
	err := captcha.WriteImage(c.Writer, captchaID, config.Admin.Captcha.Width, config.Admin.Captcha.Height)
	if err != nil {
		if err == captcha.ErrNotFound {
			ginplus.ResError(c, errors.ErrNotFound)
			return
		}
		ginplus.ResError(c, errors.WithStack(err))
		return
	}
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Content-Type", "image/png")
}
