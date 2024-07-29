package router

import (
	"github.com/gin-gonic/gin"
)

func Gin(port string) {
	router := gin.Default()

	router.POST("/hello", handler3)

	err := router.Run(port)
	if err != nil {
		panic(err)
	}
}

func handler3(c *gin.Context) {
	c.Writer.Write([]byte("Hello World!"))
}
