package logs

import "gorm.io/gorm"

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
