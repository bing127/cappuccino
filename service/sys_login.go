package service

import (
	"cappuccino/db"
	"github.com/gin-gonic/gin"
)

func Verify(ctx *gin.Context, userName, password string) (*db.SysUser, error) {
	user,err := SysUserIsExistByParams("user_name",userName)
	if err!= nil {
		return nil,err
	}
	return user,nil
}