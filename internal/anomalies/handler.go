package anomalies

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AnomalyHandler struct{
	Anomaly Anomalies
}

type Anomalies interface{
	Create(anomaly *Anomaly) error
	GetAll() ([]Anomaly, error)
	GetByID(id uint) (*Anomaly, error)
}

func NewAnomalyHandler(s Anomalies) *AnomalyHandler{
	return &AnomalyHandler {
		Anomaly: s,
	}
}

func (h *AnomalyHandler) GetAnomalys(c *gin.Context){ 
	anom, err := h.Anomaly.GetAll()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, anom)
}

func (h *AnomalyHandler) GetAnomalysID(c *gin.Context){ 
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неправильный айди"})
	}

	id := uint(id64)
	anomItem, err := h.Anomaly.GetByID(id)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, anomItem)
}
