package db

import (
	"cappuccino/utils"
	"errors"
	"github.com/jinzhu/gorm"
	"strings"
)

type SysUser struct {
	ID           string 	`gorm:"primary_key;",json:"id"`
	LoginName    string 	`gorm:"not null;unique;type:varchar(256);",json:"login_name"`
	Password     string 	`gorm:"not null;type:varchar(256);",json:"password"`
	Phone        string 	`gorm:"not null;type:varchar(16);",json:"phone"`
	Token        string 	`gorm:"type:varchar(512)",json:"token"`
	NickName   	 string 	`gorm:"not null;type:varchar(256);",json:"nick_name"`
	CreateDate   string  	`gorm:"not null;type:varchar(30);",json:"create_date"`
	CreateAt	 string 	`gorm:"not null;type:varchar(256);",json:"create_at"`
	UpdateDate   string 	`gorm:"not null;type:varchar(30);",json:"update_date"`
	UpdateAt     string 	`gorm:"not null;type:varchar(256);",json:"update_at"`
}

func (user *SysUser) Create() error {
	if utils.IsStringEmpty(user.LoginName){
		utils.PrintErr("SysUser.Create", "没有传递必要的参数")
		return errors.New("没有传递必要的参数")
	}
	if !utils.IsStringEmpty(user.Phone) {
		var count uint64
		err := getInstance().Model(&SysUser{}).Where("phone = ?", user.Phone).Count(&count).Error
		if err != nil {
			utils.PrintCallErr("User.Create", "Count Phone User", err)
			return err
		}
		if count != 0 {
			utils.PrintErr("SysUser.Create", "该手机号已存在")
			return errors.New("该手机号已存在")
		}
	}
	err := getInstance().Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			utils.PrintErr("User.Create", "用户已存在")
			return errors.New("用户已存在")
		}
		utils.PrintCallErr("User.Create", "创建用户", err)
		return err
	}
	return nil
}

func (user *SysUser) GetByLoginName() (err error) {
	if utils.IsStringEmpty(user.LoginName) {
		utils.PrintErr("User.GetByLoginName", "没有传递必要的参数")
		return errors.New("没有传递必要的参数")
	}

	err = getInstance().Where("login_name = ?", user.LoginName).Find(user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.PrintErr("User.GetByUserName", "用户不存在")
			return errors.New("用户不存在")
		}

		utils.PrintCallErr("User.GetByUserName", "User.GetByUserName", err)
		return err
	}

	return nil
}

func (user *SysUser) GetByLoginNameAndPassword() error {
	if utils.IsStringEmpty(user.LoginName) || utils.IsStringEmpty(user.Password) {
		utils.PrintErr("User.GetByUserName", "没有传递必要的参数")
		return errors.New("没有传递必要的参数")
	}
	err := getInstance().Where("login_name = ? AND password = ?", user.LoginName, user.Password).Find(user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.PrintErr("User.GetByLoginAndPassword", "用户不存在或密码错误")
			return errors.New("用户不存在或密码错误")
		}
		utils.PrintCallErr("User.GetByLoginAndPassword", "User.GetByLoginAndPassword", err)
		return err
	}

	return nil
}

func (user *SysUser) GetById() error {
	err := getInstance().Where("id = ?", user.ID).First(user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.PrintErr("User.GetById", "Find user", errors.New("用户不存在"))
			return errors.New("用户不存在")
		}
		utils.PrintCallErr("User.GetById", "Find user", err)
		return err
	}

	return nil
}

func (user *SysUser) GetByToken() error {
	if utils.IsStringEmpty(user.Token) {
		utils.PrintErr("User.GetByToken", "没有传递必要的参数")
		return errors.New("没有传递必要的参数")
	}

	err := getInstance().Where("token = ?", user.Token).First(user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.PrintErr("User.GetByToken", "Find user", errors.New("用户不存在"))
			return errors.New("用户不存在")
		}

		utils.PrintCallErr("User.GetByToken", "Find user", err)
		return err
	}

	return nil
}

func (user *SysUser) UpdateAccessToken() error {
	if len(user.ID) == 0 || utils.IsStringEmpty(user.Token) {
		utils.PrintErr("User.UpdateAccessToken", "没有传递必要的参数")
		return errors.New("没有传递必要的参数")
	}

	err := getInstance().Model(user).Update("token", user.Token).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			utils.PrintErr("User.UpdateAccessToken", errors.New("用户不存在"))
			return errors.New("用户不存在")
		}

		utils.PrintCallErr("User.UpdateAccessToken", "Find user", err)
		return err
	}

	return err
}




