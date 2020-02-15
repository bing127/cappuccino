package router

import (
	"cappuccino/api/admin"
	"cappuccino/config"
	"cappuccino/middleware"
	"github.com/gin-gonic/gin"
)

var (
	adminPostRouter = map[string]gin.HandlerFunc{
		"/user":    admin.CreateUser,
	}

	adminDeleteRouter = map[string]gin.HandlerFunc{
		//"/user/:id":  v1.DeleteGate,
	}

	adminPutRouter = map[string]gin.HandlerFunc{
		"/user": 	admin.UpdateUser,
	}

	adminGetRouter = map[string]gin.HandlerFunc{
		"/challenge":         admin.Challenge,
		//"/user/:id":                  v1.GetRoom,
	}
)

func initAdminRouter(r *gin.Engine) {
	r.Use(middleware.CorsHandler())
	r.GET(config.GetAppConfig().App.ApiPrefix+"/api/admin", admin.Admin)
	r.POST(config.GetAppConfig().App.ApiPrefix+"/api/admin/getAccessToken", admin.GetAccessToken)
	groupAdmin := r.Group(config.GetAppConfig().App.ApiPrefix+"/api/admin", middleware.JWTAuth())

	for path, f := range adminGetRouter {
		groupAdmin.GET(path, f)
	}

	for path, f := range adminPostRouter {
		groupAdmin.POST(path, f)
	}

	for path, f := range adminDeleteRouter {
		groupAdmin.DELETE(path, f)
	}

	for path, f := range adminPutRouter {
		groupAdmin.PUT(path, f)
	}
}