package db

import (
	"fmt"
	"log"

	"github.com/SergoHop/log-analyzer/internal/anomalies"
	"github.com/SergoHop/log-analyzer/internal/config"
	logs "github.com/SergoHop/log-analyzer/internal/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//чисто соединение с бд(также мы использовали AutoMigrate для миграции с наших моделей(спасибо gorm!!!))
func Init(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort,
	)


	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)

	}
	db.AutoMigrate(&logs.Log{}, &anomalies.Anomaly{})
	return db
}
 