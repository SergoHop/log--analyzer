package logs

import (
	"net/http"
	"strconv"

	"github.com/SergoHop/log-analyzer/internal/anomalies"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

type LoggerHandler struct{
	Log logger
	AnomalyRepo *anomalies.AnomalyRepository
}

type logger interface{
	CreateLogs(log *Log) error
	GetAllLogs() ([]Log, error)
	GetLogByID(id uint) (*Log, error)
	DeleteLog(id uint) (error)
	CountLogsByUserInLastMinute(userID uint) (int64, error)
}


func NewLoggerhandler(s logger, l *anomalies.AnomalyRepository) *LoggerHandler{
	return &LoggerHandler{
		Log: s,
		AnomalyRepo: l,
	}
}

func (h *LoggerHandler) CreateLogs(c *gin.Context){
	var newItem Log

	if err:= c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Log.CreateLogs(&newItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if newItem.EventType == "error"{
		anomaly := anomalies.Anomaly{
			LogID:  newItem.ID, 
              UserID: newItem.UserID,
              Rule:   "error_event",
		}
		h.AnomalyRepo.Create(&anomaly)
	}

	count, _ := h.Log.CountLogsByUserInLastMinute(newItem.UserID)
	if count > 5 {
    	h.AnomalyRepo.Create(&anomalies.Anomaly{
        	LogID: newItem.ID,
        	UserID: newItem.UserID,
        	Rule: "too_many_logs",
    	})
	}
	c.JSON(http.StatusOK, gin.H{"message": "создано"})
}

func (h *LoggerHandler) GetLogs(c *gin.Context){
	logs, err := h.Log.GetAllLogs()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logs)
}

func (h *LoggerHandler) GetLog(c *gin.Context){
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10 ,32)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "неправельный айди"})
		return
	}
	id := uint(id64)
	logItem, err := h.Log.GetLogByID(id)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, logItem)
}

func (h *LoggerHandler) DeleteLog(c *gin.Context){
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10 ,32)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"message": "лог удален"})
		return
	}

	id := uint(id64)
	if err := h.Log.DeleteLog(id); err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "лог удален"})
}