package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Println("error")
		return
	}
}
