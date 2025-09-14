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
