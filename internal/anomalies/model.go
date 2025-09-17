package anomalies

import "time"

type Anomaly struct {
    ID        uint      `gorm:"primaryKey"`
    LogID     uint     
    UserID    uint      
    Rule      string    
    CreatedAt time.Time 
}

type AnomalyStats struct {
	Total        int64 `json:"total"`
	ErrorEvents  int64 `json:"error_events"`
	TooManyLogs  int64 `json:"too_many_logs"`
}