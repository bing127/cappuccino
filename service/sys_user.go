package service

import (
	"cappuccino/db"
	"cappuccino/utils"
	"errors"
	"github.com/jinzhu/gorm"
)

// 判断登陆账号是否存在
func LoginNameIsExist(name string) bool {
	user := db.SysUser{LoginName:name}
	err := user.GetByLoginName()
	if err != nil {
		return true
	}
	return false
}

//根据用户Id判断该用户是否存在
func UserIsExist(id string) bool {
	user := db.SysUser{ID:id}
	err := user.GetById()
	if err != nil {
		return true
	}
	return false
}

//更新用户信息
func UpdateUserInfo(user *db.SysUser,id string) error {
	err := db.GormDb.Model(&user).Where("id = ? ",id).Updates(user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.PrintErr("User.UpdateUserInfo", "Find user", errors.New("用户不存在"))
			return errors.New("用户不存在")
		}

		utils.PrintCallErr("User.UpdateUserInfo", "Find user", err)
		return err
	}
	return nil
}