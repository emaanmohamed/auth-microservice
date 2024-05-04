package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func Services() gin.HandlerFunc {
	return func(g *gin.Context) {
		segments := strings.Split(g.Request.RequestURI, "/")
		log.Fatal(segments, len(segments))
		println(segments)
	}
}

func ServiceList() map[string]bool {

}
