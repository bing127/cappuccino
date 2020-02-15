package admin

import (
	"cappuccino/config"
	"cappuccino/utils/apiRequest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Admin godoc
// @Summary 获取后台版本
// @Description 获取后台版本
// @Tags 后台
// @Success 200 { string } json
// @Router /cappuccino/api/admin [get]
func Admin(c *gin.Context) {
	result := make(map[string]interface{})
	result["data"] = config.GetAppConfig().Version
	c.JSON(http.StatusOK, apiRequest.ResponseJson("操作成功",result,true))
}

