package v1

import (
	"cappuccino/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// V1 godoc
// @Summary 获取平台版本
// @Description 获取平台版本
// @Tags 获取平台版本
// @Success 200 {string} string
// @Router /cappuccino/api/v1 [get]
func V1(c *gin.Context) {
	result := make(map[string]interface{})
	result["date"] = utils.GetCurrentTimeStamp()
	c.JSON(http.StatusOK,utils.ResponseJson("请求成功",result,true,""))
}
