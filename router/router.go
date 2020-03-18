package router

import (
	"cappuccino/config"
	_ "cappuccino/docs"
	"cappuccino/middleware"
	"cappuccino/utils"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"os"
)

func InitRouter(r *gin.Engine) {
	if !utils.PathExists(config.Admin.File.Dir) {
		err := os.MkdirAll(config.Admin.File.Dir, os.ModePerm|os.ModeDir)
		if err != nil {
			panic(err)
		}
	}
	// 静态资源路径
	r.Static(config.Admin.File.Path, config.Admin.File.Dir)

	// 模板视图
	//r.LoadHTMLGlob("view/*")
	r.LoadHTMLFiles("view/admin/index.html")

	// 处理跨域请求
	r.Use(middleware.CorsHandler())

	// 从panic中恢复
	r.Use(gin.Recovery())

	// swagger
	r.GET(config.Admin.App.ApiPrefix+"/doc/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 404路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, utils.NotFound())
	})
	// 管理后端使用的API
	initAdminRouter(r)
}
