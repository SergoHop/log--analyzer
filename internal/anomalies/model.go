package anomalies

import "time"

type Anomaly struct {
    ID        uint      `gorm:"primaryKey"`
    LogID     uint     
    UserID    uint      
    Rule      string    
    CreatedAt time.Time 
}

