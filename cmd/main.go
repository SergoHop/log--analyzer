package main

import (
	"github.com/SergoHop/log-analyzer/internal/database"
	"github.com/SergoHop/log-analyzer/internal/logs"
	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()
	
	repo := logs.NewLogRepository(DB)
	
	handler := logs.NewLoggerhandler(repo)

	r := gin.Default()
    r.POST("/logs", handler.CreateLogs)
    r.GET("/logs", handler.GetLogs)
    r.GET("/logs/:id", handler.GetLog)
    r.DELETE("/logs/:id", handler.DeleteLog)
	r.Run(":8080")
}