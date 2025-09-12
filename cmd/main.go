package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SergoHop/log-analyzer/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	DB := db.Init()
	fmt.Println("бд норм", DB != nil)
	log.Println("бд конект")

	r := gin.Default()
	

	r.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "все норм")
	})
	r.Run(":8080")
}