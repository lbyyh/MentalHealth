package app

import (
	"MentalHealth-Platform/app/model"
	"MentalHealth-Platform/app/router"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	model.InitConfig()

	// 初始化Gin引擎
	r := gin.Default()

	// 初始化路由
	router.New()

	// 设置 tools 包中的 Redis 客户端
	//tools.SetRedisClient(model.RedisDB)

	// 创建自定义的 HTTP 服务器
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// 启动服务器并在新的 goroutine 中运行
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	fmt.Println("Service started successfully.")

	// 优雅停止服务的逻辑
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // 阻塞直至接收到终止信号
	fmt.Println("Shutting down server...")

	// 创建一个超时时间为5秒的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
