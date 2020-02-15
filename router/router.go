package router

import (
	"cappuccino/config"
	_ "cappuccino/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter(r *gin.Engine) {
	// 静态资源路径
	r.Static(config.GetAppConfig().Static.Url, config.GetAppConfig().Static.Dir)
	// swagger
	r.GET(config.GetAppConfig().App.ApiPrefix + "/api/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 管理后端使用的API
	initAdminRouter(r)
}
