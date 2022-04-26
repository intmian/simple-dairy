package main

import "github.com/gin-gonic/gin"

func addRoute(g *gin.Engine) {
	g.GET("/")
}