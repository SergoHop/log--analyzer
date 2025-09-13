package logs

import (
	"time"

	"gorm.io/gorm"
)

type Log struct{
	gorm.Model
	UserID uint
	EventType string
	EventTime time.Time
	Metadata string
}