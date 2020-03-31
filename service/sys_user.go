package service

import (
	"cappuccino/db"
	"errors"
	"github.com/jinzhu/gorm"
)

func SysUserIsExistByParams(params map[string]interface{}) (*db.SysUser,error) {
	user := db.SysUser{}
	err := db.GetInstance().Where(params).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil,errors.New("未找到符合条件信息")
		}
		return nil,errors.New("查询失败")
	}
	return &user,nil
}
