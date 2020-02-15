package service

import (
	"cappuccino/middleware"
	"cappuccino/utils"
	"errors"
	"cappuccino/db"
)

func GetAndUpdateAccessToken(loginName,password string) (string, error) {
	if utils.IsStringEmpty(loginName) || utils.IsStringEmpty(password) {
		utils.PrintErr("GetAndUpdateAccessToken", "没有传递必要的参数")
		return "", errors.New("没有传递必要的参数")
	}

	user := &db.SysUser{LoginName: loginName, Password: utils.GetMd5(password)}
	err := user.GetByLoginNameAndPassword()
	if err != nil {
		utils.PrintCallErr("GetAndUpdateAccessToken", "user.GetByUserNameAndPassword", err)
		return "", err
	}

	token,err := middleware.GenerateToken(user)
	if err!=nil {
		utils.PrintCallErr("GetAccessToken", "user.GetUserAccessToken", err)
		return "", err
	}

	user.Token = token
	err = user.UpdateAccessToken()
	if err != nil {
		utils.PrintCallErr("GetAccessToken", "user.UpdateAccessToken", err)
		return "", err
	}

	return token, nil
}
