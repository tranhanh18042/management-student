package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/aciiJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Trần Đức Hạnh",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
