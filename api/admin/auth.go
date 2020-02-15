package admin

import (
	"cappuccino/api/dto"
	"cappuccino/middleware"
	"cappuccino/service"
	"cappuccino/utils/apiRequest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Challenge godoc
// @Summary Token检查
// @Description Token检查
// @Tags 后台
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Success 200 { string } json
// @Failure 400 { string } json
// @Failure 401 { string } json
// @Failure 500 { string } json
// @Router /cappuccino/api/admin/challenge [get]
func Challenge(c *gin.Context) {
	result := make(map[string]interface{})
	result["data"] = middleware.GetUserInfoByToken(c)
	c.JSON(http.StatusOK,apiRequest.ResponseJson("成功",result,true))
}

// GetAccessToken godoc
// @Summary 获取Token
// @Description 获取Token
// @Tags 后台
// @Accept  json
// @Produce json
// @Param GetAccessTokenRequest body dto.GetAccessTokenRequest true "获取Token信息"
// @Success 200 {object} dto.GetAccessTokenResponse
// @Failure 400 {object} dto.GetAccessTokenResponse
// @Failure 401 {object} dto.GetAccessTokenResponse
// @Failure 500 {object} dto.GetAccessTokenResponse
// @Router /cappuccino/api/admin/getAccessToken [post]
func GetAccessToken(c *gin.Context) {
	result := make(map[string]interface{})
	request := &dto.GetAccessTokenRequest{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiRequest.ResponseJson(err.Error(),nil,true))
		return
	}

	token, err := service.GetAndUpdateAccessToken(request.LoginName, request.Password)
	if err != nil {
		c.JSON(http.StatusOK, apiRequest.ResponseJson(err.Error(),nil,true))
		return
	}
	result["token"] = token
	c.JSON(http.StatusOK, apiRequest.ResponseJson("获取Token成功",result,true))
}
