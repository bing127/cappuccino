package service

import (
	"cappuccino/db"
	"fmt"
	"github.com/jinzhu/gorm"
	"errors"
)

func SysUserIsExistByParams(key,value string) (*db.SysUser,error) {
	keyStr := fmt.Sprintf("%s=?",key)
	user := db.SysUser{}
	err := db.GetInstance().Where(keyStr,value).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil,errors.New("未找到符合条件信息")
		}
		return nil,errors.New("查询失败")
	}
	return &user,nil
}
