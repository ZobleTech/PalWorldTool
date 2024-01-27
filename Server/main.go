package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/Routes"
)

func Server() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	basic := router.Group("/basic")
	{
		basic.POST("/create_ini", Routes.CreateInI)
		basic.POST("/analyze_ini", Routes.AnalyzeIni)
	}
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	Server()
}
