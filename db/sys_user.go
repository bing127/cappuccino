package db

import "time"

// User 用户实体
type SysUser struct {
	ID         string     `gorm:"column:id;primary_key;size:64;index;"` // 记录内码
	UserName   *string    `gorm:"column:user_name;size:64;index;"`      // 用户名
	RealName   *string    `gorm:"column:real_name;size:64;index;"`      // 真实姓名
	Password   *string    `gorm:"column:password;size:40;"`             // 密码(sha1(md5(明文))加密)
	Email      *string    `gorm:"column:email;size:255;index;"`         // 邮箱
	Phone      *string    `gorm:"column:phone;size:20;index;"`          // 手机号
	Status     *int       `gorm:"column:status;index;"`                 // 状态(1:启用 2:停用)
	CreateBy   *string    `gorm:"column:create_by;size:64;"`            // 创建者
	CreateDate *string    `gorm:"column:create_date;size:64;"`          // 创建时间
	UpdateBy   *time.Time `gorm:"column:update_by;"`                    // 更新者
	UpdateDate *time.Time `gorm:"column:update_date;"`                  // 更新时间
}

// TableName 表名
func (a SysUser) TableName() string {
	return a.Model.TableName("sys_user")
}

// UserRole 用户角色关联实体
type SysUserRole struct {
	UserID string `gorm:"column:user_id;size:36;index;"` // 用户内码
	RoleID string `gorm:"column:role_id;size:36;index;"` // 角色内码
	Model
}

// TableName 表名
func (a SysUserRole) TableName() string {
	return a.Model.TableName("sys_user_role")
}
