package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func Services() gin.HandlerFunc {
	return func(g *gin.Context) {
		segments := strings.Split(g.Request.RequestURI, "/")
		segmentOne := strings.Split(segments[1], "?")
		services := ServiceList()
		if ok := services[segmentOne[1]]; ok == "" {
			g.JSON(401, gin.H{
				"message": "Service not found",
				"status":  "error",
				"code":    401,
			})
			g.Abort()
			return
		}
		serviceUrl := services[segmentOne[0]]
		log.Fatal(serviceUrl)

		g.JSON(200, gin.H{
			"message": "Service found",
			"status":  "ok",
			"code":    200,
		})
	}
}

func ServiceList() map[string]string {
	m := make(map[string]string)
	m["users"] = "http://localhost:8082"
	m["posts"] = "http://localhost:8083"
	return m
}
