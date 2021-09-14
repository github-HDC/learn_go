package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func handle(c ...interface{}) {
	fmt.Println(c)
}
func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
