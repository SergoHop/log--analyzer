package main

import (
	"github.com/SergoHop/log-analyzer/internal/anomalies"
	"github.com/SergoHop/log-analyzer/internal/database"
	"github.com/SergoHop/log-analyzer/internal/logs"
	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()
	
	logrepo := logs.NewLogRepository(DB)
	anomalyRepo := anomalies.NewAnomalyRepository(DB)
	handler := logs.NewLoggerhandler(logrepo,anomalyRepo)
	anomalyHandler := anomalies.NewAnomalyHandler(anomalyRepo)


	r := gin.Default()
	//логи
    r.POST("/logs", handler.CreateLogs)
    r.GET("/logs", handler.GetLogs)
    r.GET("/logs/:id", handler.GetLog)
    r.DELETE("/logs/:id", handler.DeleteLog)
	//аномалии
	r.GET("/anomaly", anomalyHandler.GetAnomalys)
	r.GET("/anomaly/:id", anomalyHandler.GetAnomalysID)
	r.Run(":8080")
}