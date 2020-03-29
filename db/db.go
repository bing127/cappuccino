package db

import (
	"cappuccino/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var gormDb *gorm.DB

// 表名前缀
var tablePrefix string

// SetTablePrefix 设定表名前缀
func SetTablePrefix(prefix string) {
	tablePrefix = prefix
}

// GetTablePrefix 获取表名前缀
func GetTablePrefix() string {
	return tablePrefix
}

// Open 打开数据库
func Open() {
	mysqlConfig := config.Admin.Db
	template := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	connStr := fmt.Sprintf(template, mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Schema)
	openDb, err := gorm.Open("mysql", connStr)
	if err != nil {
		log.Println(err.Error())
		panic("数据库连接异常:**********************" + mysqlConfig.Host + "**********************")
	}

	openDb.LogMode(true)
	openDb.SingularTable(true)
	gormDb = openDb

	autoMigrate(&SysUser{}, &SysUserRole{}, &SysRole{}, &SysRoleMenu{}, &SysMenu{}, &SysMenuAction{}, &SysMenuResource{})
}

// Close 关闭数据库
func Close() {
	_ = gormDb.Close()
	gormDb = nil
}

// GetInstance 获取数据库实例
func GetInstance() *gorm.DB {
	return gormDb
}

func autoMigrate(values ...interface{}) {
	gormDb.AutoMigrate(values...)
}

// Model base model
type Model struct {
	CreatedAt time.Time  `gorm:"column:created_at;"`
	UpdatedAt time.Time  `gorm:"column:updated_at;"`
	DeletedAt *time.Time `gorm:"column:deleted_at;index;"`
}

// TableName table name
func (Model) TableName(name string) string {
	return fmt.Sprintf("%s%s", GetTablePrefix(), name)
}
