package db

import (
	"cappuccino/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var GormDb *gorm.DB


// Open 打开数据库
func Open() {
	mysqlConfig := config.GetAppConfig().Database
	template := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	connStr := fmt.Sprintf(template, mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Address, mysqlConfig.Schema)
	openDb, err := gorm.Open("mysql", connStr)
	if err != nil {
		log.Println(err.Error())
		panic("数据库连接异常")
	}

	openDb.LogMode(config.GetAppConfig().Database.OpenLog)
	//全局设置表名不可以为复数形式。
	openDb.SingularTable(true)

	GormDb = openDb

	autoMigrate(&SysUser{})
}

// Close 关闭数据库
func Close() {
	_ = GormDb.Close()
	GormDb = nil
}

// GetInstance 获取数据库实例
func getInstance() *gorm.DB {
	return GormDb
}

func autoMigrate(values ...interface{}) {
	GormDb.AutoMigrate(values...)
}
