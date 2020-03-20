package router

import (
	"cappuccino/config"
	v1 "cappuccino/controller/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	v1PostRouter = map[string]gin.HandlerFunc{
	}

	v1DeleteRouter = map[string]gin.HandlerFunc{
	}

	v1PutRouter = map[string]gin.HandlerFunc{
	}

	v1GetRouter = map[string]gin.HandlerFunc{
	}
)

func initAdminRouter(r *gin.Engine) {
	// 后台管理系统
	r.GET(config.Admin.App.ApiPrefix+"/admin", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title":config.Admin.App.Name,
		})
	})
	r.GET(config.Admin.App.ApiPrefix+"/api/v1", v1.V1)
	//r.POST(config.Admin.App.ApiPrefix+"/api/admin/getAccessToken", v1.GetAccessToken)

	groupV1 := r.Group(config.Admin.App.ApiPrefix + "/api/v1")

	for path, f := range v1GetRouter {
		groupV1.GET(path, f)
	}

	for path, f := range v1PostRouter {
		groupV1.POST(path, f)
	}

	for path, f := range v1DeleteRouter {
		groupV1.DELETE(path, f)
	}

	for path, f := range v1PutRouter {
		groupV1.PUT(path, f)
	}
}