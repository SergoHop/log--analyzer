package main

import (
	"fmt"
	"log"
	"github.com/SergoHop/log-analyzer/internal/database"
)

func main() {
	DB := db.Init()
	fmt.Println("бд норм", DB != nil)
	log.Println("бд конект")
}