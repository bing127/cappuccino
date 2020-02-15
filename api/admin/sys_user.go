package admin

import (
	"cappuccino/api/dto"
	"cappuccino/db"
	"cappuccino/middleware"
	"cappuccino/service"
	"cappuccino/utils"
	"cappuccino/utils/apiRequest"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser godoc
// @Summary 创建用户
// @Description 创建用户
// @Tags 后台
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Param CreateUserRequest body dto.CreateUserRequest true "用户信息"
// @Success 200 { string } json
// @Failure 400 { string } json
// @Failure 500 { string } json
// @Router /cappuccino/api/admin/user [post]
func CreateUser(c *gin.Context) {
	request := &dto.CreateUserRequest{}
	result := make(map[string]interface{})
	err := c.ShouldBindJSON(request)
	if err != nil {
		c.JSON(http.StatusOK, apiRequest.ResponseJson("创建用户失败",nil,true))
		return
	}

	userIsExist := service.LoginNameIsExist(request.LoginName)
	if !userIsExist {
		c.JSON(http.StatusOK, apiRequest.ResponseJson("用户已存在",nil,true))
		return
	}
	uId := utils.GetUUID()
	user := &db.SysUser{
		LoginName:request.LoginName,
		NickName:request.NickName,
		ID:uId,
		Password:utils.GetMd5("123456"),
		CreateDate:utils.GetCurrentTime(),
		CreateAt:middleware.GetUserInfoByToken(c).ID,
		UpdateDate:utils.GetCurrentTime(),
		UpdateAt:middleware.GetUserInfoByToken(c).ID,
	}
	userErr := user.Create()
	if userErr != nil {
		c.JSON(http.StatusOK, apiRequest.ResponseJson(userErr.Error(),nil,true))
		return
	}
	result["userId"] = uId
	c.JSON(http.StatusOK, apiRequest.ResponseJson("创建用户成功",result,true))
}

// CreateUser godoc
// @Summary 更新用户
// @Description 更新用户
// @Tags 后台
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Param UpdateUserRequest body dto.UpdateUserRequest true "用户信息"
// @Success 200 { string } json
// @Failure 400 { string } json
// @Failure 500 { string } json
// @Router /cappuccino/api/admin/user [put]
func UpdateUser(c *gin.Context)  {
	request := &dto.UpdateUserRequest{}
	//result := make(map[string]interface{})
	err := c.ShouldBindJSON(request)
	if err != nil {
		c.JSON(http.StatusOK, apiRequest.ResponseJson("缺少提交必要参数",nil,true))
		return
	}

	userIsExist := service.UserIsExist(request.ID)
	if userIsExist {
		c.JSON(http.StatusOK, apiRequest.ResponseJson("用户不存在",nil,true))
		return
	}
	user := &db.SysUser{
		Phone:      request.Phone,
		NickName:   request.NickName,
		UpdateDate: utils.GetCurrentTime(),
		UpdateAt:   middleware.GetUserInfoByToken(c).ID,
	}
	updateErr := service.UpdateUserInfo(user,request.ID)
	if updateErr != nil {
		c.JSON(http.StatusOK, apiRequest.ResponseJson("更新失败",nil,true))
		return
	}
	c.JSON(http.StatusOK, apiRequest.ResponseJson("用户更新成功",nil,true))
	return

}

// CreateUser godoc
// @Summary 删除用户
// @Description 删除用户
// @Tags 后台
// @Accept  json
// @Produce json
// @Param Authorization header string true "Authentication header"
// @Param CreateUserResponse body dto.CreateUserResponse true "用户信息"
// @Success 200 { string } json
// @Failure 400 { string } json
// @Failure 500 { string } json
// @Router /cappuccino/api/admin/user [delete]
func DeleteUser(c *gin.Context)  {

}