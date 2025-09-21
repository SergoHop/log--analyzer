package anomalies

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type AnomalyRepository struct{
	Db *gorm.DB
}

func NewAnomalyRepository(Db *gorm.DB) *AnomalyRepository{
	return &AnomalyRepository{Db: Db}
}

func (r *AnomalyRepository) Create(anomaly *Anomaly) error{
	result := r.Db.Create(anomaly)
	return result.Error
}

func (r *AnomalyRepository) GetAll() ([]Anomaly, error){
	var anomalies []Anomaly
	result := r.Db.Find(&anomalies)
	return anomalies, result.Error
}

func (r *AnomalyRepository) GetByID(id uint) (*Anomaly, error){
	var anomalie Anomaly
	result := r.Db.Where("id = ?", id).First(&anomalie)
	return &anomalie, result.Error
}
//нужно для статистики, то есть юзер выдал 4 ошибки, ты кидаешь стат и вот он как на ладони все видно, скок и че по чем
func (r *AnomalyRepository) GetStats(userID *uint) (AnomalyStats, error){
	var stats AnomalyStats
	q := r.Db.Model(&Anomaly{})
	if userID != nil{
		q  = q.Where("user_id = ?", *userID)
	}
	if err := q.Count(&stats.Total).Error; err != nil{
		return stats, err
	}

	q2 := r.Db.Model(&Anomaly{})
    if userID != nil {
        q2 = q2.Where("user_id = ?", *userID)
    }
    if err := q2.Where("rule = ?", "error_event").Count(&stats.ErrorEvents).Error; err != nil {
        return stats, err
    }

    q3 := r.Db.Model(&Anomaly{})
    if userID != nil {
        q3 = q3.Where("user_id = ?", *userID)
    }
    if err := q3.Where("rule = ?", "too_many_logs").Count(&stats.TooManyLogs).Error; err != nil {
        return stats, err
    }
	return stats, nil
}

func (r *AnomalyRepository) GetRecentByRule(userID uint, rule string, since time.Time) ([]Anomaly, error) {
	var anomalies []Anomaly
	result := r.Db.Where("user_id = ? AND rule = ? AND created_at >= ?", userID, rule, since).Find(&anomalies)
	return anomalies, result.Error
}

func (r *AnomalyRepository) GetLastTooManyLogs(userID uint) (*Anomaly, error) {
	var anomaly Anomaly
	result := r.Db.Where("user_id = ? AND rule = ?", userID, "too_many_logs").
		Order("created_at desc").
		First(&anomaly)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // нет записей — можно создавать
		}
		return nil, result.Error
	}
	return &anomaly, nil
}
