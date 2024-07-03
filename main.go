package main

import (
	"Calendar/handlers"
	"Calendar/scheduler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 路由转发
	handlers.RegisterRoutes(router)

	// 启动异步线程进行不断监听到期的提醒
	go scheduler.StartReminderChecker()

	router.Run(":8080")
}
