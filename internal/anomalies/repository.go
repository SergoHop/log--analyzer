package anomalies

import "gorm.io/gorm"

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
	if err := q.Where("rule = ?", "error_event").Count(&stats.ErrorEvents).Error; err != nil{
		return stats, err
	}
	if err := q.Where("rule = ?", "too_many_logs").Count(&stats.TooManyLogs).Error; err != nil{
		return stats, err
	}
	return stats, nil
}