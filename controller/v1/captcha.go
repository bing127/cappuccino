package v1

import (
	"cappuccino/config"
	"cappuccino/utils"
	"github.com/LyricTian/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCaptcha 获取图形验证码信息
func GetCaptcha(ctx *gin.Context) {
	captchaID := captcha.NewLen(config.Admin.Captcha.Length)
	result := make(map[string]interface{})

	err := captcha.WriteImage(ctx.Writer, captchaID, config.Admin.Captcha.Width, config.Admin.Captcha.Height)
	if err != nil {
		if err == captcha.ErrNotFound {
			result["captcha"] = ""
		}
		result["captcha"] = ""
	}
	ctx.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Writer.Header().Set("Pragma", "no-cache")
	ctx.Writer.Header().Set("Expires", "0")
	ctx.Writer.Header().Set("Content-Type", "image/png")
	ctx.JSON(http.StatusOK,utils.ResponseJson("请求成功",result,true,""))
}
