package service

import (
	"cappuccino/db"
	"cappuccino/errors"
	"cappuccino/utils"
	"github.com/gin-gonic/gin"
)

func Verify(ctx *gin.Context, userName, password string) (*db.SysUser, error) {
	params := make(map[string]interface{})
	params["user_name"] = userName
	user,err := SysUserIsExistByParams(params)
	if err!= nil {
		return nil,errors.New("用户不存在")
	}
	params["password"] = utils.Hmac(password)
	user,err = SysUserIsExistByParams(params)
	if err!= nil {
		return nil,errors.New("密码错误")
	}
	return user,nil
}