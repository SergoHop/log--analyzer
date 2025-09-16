package logs

import (
	"time"

	"gorm.io/gorm"
)

type LogRepository struct{
	Db *gorm.DB
}

func NewLogRepository(Db *gorm.DB) *LogRepository{
	return &LogRepository{Db: Db}
}

func (r *LogRepository) GetAllLogs() ([]Log, error){
	var logs []Log
	result := r.Db.Find(&logs)
	return logs, result.Error
}

func (r *LogRepository) GetLogByID(id uint) (*Log, error){
	var log Log
	result := r.Db.Where("id = ?", id).First(&log)
	return &log, result.Error
}

func (r *LogRepository) CreateLogs(log *Log) error{
	result := r.Db.Create(log)
	return result.Error
}

func (r *LogRepository) DeleteLog(id uint) error{
	result := r.Db.Where("id = ?", id).Delete(&Log{})
	return result.Error
}

func (r *LogRepository) CountLogsByUserInLastMinute(userID uint) (int64, error){
	var count int64
	r.Db.Model(&Log{}).Where("user_id = ? AND created_at >= ?", userID, time.Now().Add(-1*time.Minute)).Count(&count)
	return count, nil
}


//это для фильтрации, тоесть ты пишель id и он выдает все логи по этому id или по error тоесть выдаст только error
func (r *LogRepository) GetFilteredLogs(userID *uint, eventType string) ([]Log, error){
	var logs []Log
	q := r.Db.Model(&Log{})

	if userID != nil{
		q  = q.Where("user_id = ?", *userID)
	}
	if eventType != ""{
		q  = q.Where("event_type = ?", eventType)
	}
	result := q.Find(&logs)
	return logs, result.Error
}
