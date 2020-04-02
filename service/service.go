package service

import (
	"cappuccino/db"
	"cappuccino/pkg"
)

// Init 初始化service
func Init() {
	// 数据库打开
	db.Open()
	// redis打开
	pkg.InitRedis()
}

// Destroy 销毁service
func Destroy() {
	db.Close()
	pkg.Close()
}