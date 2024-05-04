package main

import (
	"github.com/gin-gonic/gin"
)

func Welcome(g *gin.Context) {
	g.JSON(200, gin.H{
		"message": "Welcome to the server",
		"status":  "ok",
		"code":    200,
	})
}

func main() {
	connect()
	Migrate()
	g := gin.Default()
	g.Use(Services())

	g.GET("/ping", Welcome)
	g.Run(":8083")

}
