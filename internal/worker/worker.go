package worker

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SergoHop/log-analyzer/internal/anomalies"
	"github.com/SergoHop/log-analyzer/internal/logs"
)

type Worker struct{
	Log logger
	Anom Anomalies
}

type logger interface{
	CreateLogs(log *logs.Log) error
	GetAllLogs() ([]logs.Log, error)
	GetLogByID(id uint) (*logs.Log, error)
	DeleteLog(id uint) (error)
	CountLogsByUserInLastMinute(userID uint) (int64, error)
	GetFilteredLogs(userID *uint, eventType string) ([]logs.Log, error)
	GetLogsSince(t time.Time) ([]logs.Log,error)
}

type Anomalies interface{
	Create(anomaly *anomalies.Anomaly) error
	GetAll() ([]anomalies.Anomaly, error)
	GetByID(id uint) (*anomalies.Anomaly, error)
	GetStats(userID *uint) (anomalies.AnomalyStats, error)
	GetRecentByRule(userID uint, rule string, since time.Time) ([]anomalies.Anomaly, error)
	GetLastTooManyLogs(userID uint) (*anomalies.Anomaly, error)
}

func NewWorker(s Anomalies, l logger) Worker{
	return Worker{
		Anom: s,
		Log: l,
	}
}

func (wp Worker) Stats(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			log.Println("Запуск воркера: проверка логов")
			from := time.Now().Add(-10 * time.Second)
			logs, err := wp.Log.GetLogsSince(from)
			if err != nil {
				log.Printf("ошибка получения логов: %v", err)
				continue
			}

			for _, logItem := range logs {
				// Правило error_event
				if logItem.EventType == "error" {
					anomaly := anomalies.Anomaly{
						LogID:  logItem.ID,
						UserID: logItem.UserID,
						Rule:   "error_event",
					}
					wp.Anom.Create(&anomaly)
				}

				// Правило too_many_logs
				count, _ := wp.Log.CountLogsByUserInLastMinute(logItem.UserID)
				if count > 5 {
					last, err := wp.Anom.GetLastTooManyLogs(logItem.UserID)
					if err != nil {
						log.Printf("ошибка получения последней аномалии: %v", err)
						continue
					}

					if last == nil || time.Since(last.CreatedAt) > time.Minute {
						wp.Anom.Create(&anomalies.Anomaly{
							LogID:  logItem.ID,
							UserID: logItem.UserID,
							Rule:   "too_many_logs",
						})
						log.Printf("Пользователь %d: найдено %d логов за минуту", logItem.UserID, count)
					} else {
						log.Printf("Пользователь %d: пропущена аномалия, т.к. была недавно", logItem.UserID)
					}
				}
			}

		case <-ctx.Done():
			fmt.Printf("остановка воркера: %v\n", ctx.Err())
			return
		}
	}
}


///