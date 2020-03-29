package db

// Menu 菜单实体
type SysMenu struct {
	ID         string  `gorm:"column:id;primary_key;size:64;index;"`           // 记录内码
	Name       *string `gorm:"column:name;size:50;index;"`         // 菜单名称
	Sort       *int    `gorm:"column:sort;index;"`                 // 排序值
	Icon       *string `gorm:"column:icon;size:255;"`              // 菜单图标
	Router     *string `gorm:"column:router;size:255;"`            // 访问路由
	Hidden     *int    `gorm:"column:hidden;index;"`               // 隐藏菜单(0:不隐藏 1:隐藏)
	ParentID   *string `gorm:"column:parent_id;size:36;index;"`    // 父级内码
	ParentPath *string `gorm:"column:parent_path;size:518;index;"` // 父级路径
	Creator    *string `gorm:"column:creator;size:36;"`            // 创建人
	Model
}

// TableName 表名
func (a SysMenu) TableName() string {
	return a.Model.TableName("sys_menu")
}

// MenuAction 菜单动作关联实体
type SysMenuAction struct {
	MenuID string `gorm:"column:menu_id;size:36;index;"` // 菜单ID
	Code   string `gorm:"column:code;size:50;index;"`    // 动作编号
	Name   string `gorm:"column:name;size:50;"`          // 动作名称
	Model
}

// TableName 表名
func (a SysMenuAction) TableName() string {
	return a.Model.TableName("sys_menu_action")
}

// MenuResource 菜单资源关联实体
type SysMenuResource struct {
	MenuID string `gorm:"column:menu_id;size:36;index;"` // 菜单ID
	Code   string `gorm:"column:code;size:50;index;"`    // 资源编号
	Name   string `gorm:"column:name;size:50;"`          // 资源名称
	Method string `gorm:"column:method;size:50;"`        // 请求方式
	Path   string `gorm:"column:path;size:255;"`         // 请求路径
	Model
}

// TableName 表名
func (a SysMenuResource) TableName() string {
	return a.Model.TableName("sys_menu_resource")
}
