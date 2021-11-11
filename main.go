package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		fmt.Println("Request started: wait 40 seconds to proceed...")
		time.Sleep(40 * time.Second)
		fmt.Println("Request done!")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("health-check heartbeat")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
