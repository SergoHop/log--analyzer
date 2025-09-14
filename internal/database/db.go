package db

import (
	"log"

	"github.com/SergoHop/log-analyzer/internal/anomalies"
	logs "github.com/SergoHop/log-analyzer/internal/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//чисто соединение с бд(также мы использовали AutoMigrate для миграции с наших моделей(спасибо gorm!!!))
func Init() *gorm.DB {
	dbURL := "postgres://sergo_user:13410285@localhost:5433/log-analyzer"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)

	}
	db.AutoMigrate(&logs.Log{}, &anomalies.Anomaly{})
	return db
}
 