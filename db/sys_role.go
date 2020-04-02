package db

import "time"

// Role 角色实体
type SysRole struct {
	ID         string     `gorm:"column:id;primary_key;size:64;index;"` // 记录内码
	Name       *string    `gorm:"column:name;size:100;index;"`          // 角色名称
	Sort       *int       `gorm:"column:sort;index;"`                   // 排序值
	Memo       *string    `gorm:"column:memo;size:200;"`                // 备注
	CreateBy   *string    `gorm:"column:create_by;size:64;"`            // 创建者
	CreateDate *time.Time    `gorm:"column:create_date;size:64;"`          // 创建时间
	UpdateBy   *string `gorm:"column:update_by;"`                    // 更新者
	UpdateDate *time.Time `gorm:"column:update_date;"`                  // 更新时间
}

// RoleMenu 角色菜单关联实体
type SysRoleMenu struct {
	RoleID     string     `gorm:"column:role_id;size:36;index;"` // 角色内码
	MenuID     string     `gorm:"column:menu_id;size:36;index;"` // 菜单内码
	Action     *string    `gorm:"column:action;size:2048;"`      // 动作权限(多个以英文逗号分隔)
	Resource   *string    `gorm:"column:resource;size:2048;"`    // 资源权限(多个以英文逗号分隔)
	CreateBy   *string    `gorm:"column:create_by;size:64;"`            // 创建者
	CreateDate *time.Time    `gorm:"column:create_date;size:64;"`          // 创建时间
	UpdateBy   *string `gorm:"column:update_by;"`                    // 更新者
	UpdateDate *time.Time `gorm:"column:update_date;"`                  // 更新时间
}
