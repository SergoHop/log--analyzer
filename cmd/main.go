package main

import (
	"context"
	"os/signal"
	"syscall"
	"os"

	"github.com/SergoHop/log-analyzer/internal/anomalies"
	"github.com/SergoHop/log-analyzer/internal/database"
	"github.com/SergoHop/log-analyzer/internal/logs"
	"github.com/SergoHop/log-analyzer/internal/worker"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	DB := db.Init()
	
	logrepo := logs.NewLogRepository(DB)
	anomalyRepo := anomalies.NewAnomalyRepository(DB)
	handler := logs.NewLoggerhandler(logrepo,anomalyRepo)
	anomalyHandler := anomalies.NewAnomalyHandler(anomalyRepo)

	wrk := worker.NewWorker(anomalyRepo, logrepo)

	go wrk.Stats(ctx)

	r := gin.Default()
	//логи
    r.POST("/logs", handler.CreateLogs)
    r.GET("/logs", handler.GetLogs)
    r.GET("/logs/:id", handler.GetLog)
    r.DELETE("/logs/:id", handler.DeleteLog)
	//аномалии
	r.GET("/anomaly", anomalyHandler.GetAnomalys)
	r.GET("/anomaly/:id", anomalyHandler.GetAnomalysID)
	//статистика для аномалий
	r.GET("/anomaly/stats", anomalyHandler.GetStats)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	//для запуска в отдельной гарутине 
	go func() {
		if err := r.Run(":8080"); err != nil{
			panic(err)
		}
	}()
	<-quit
	cancel()

}