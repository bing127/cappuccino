package main

import (
	"cappuccino/config"
	"cappuccino/router"
	"cappuccino/service"
	"context"
	"github.com/gin-gonic/gin"
	DEATH "gopkg.in/vrecan/death.v3"
	"log"
	"net/http"
	"syscall"
	"time"
)

// @title cappuccino
// @version latest
// @BasePath /
func main() {
	service.Init()
	defer service.Destroy()

	// 初始化图形验证码
	service.InitCaptcha()
	// 禁用控制台颜色
	//gin.DisableConsoleColor()

	// 创建记录日志的文件
	//f, _ := os.Create(config.Admin.App.Name + ".log")
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	gin.SetMode(config.Admin.Server.RunMode)
	r := gin.Default()

	router.InitRouter(r)

	srv := &http.Server{
		Addr:    ":" + config.Admin.Server.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	death := DEATH.NewDeath(syscall.SIGINT, syscall.SIGTERM)
	_ = death.WaitForDeath()
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
