package main

import (
	"First_program/database"
	. "First_program/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	defer database.SQLdb.Close()
	router.GET("/", IndexAPI)
	router.POST("/article/add", AddarticleAPI)
	router.POST("/article/searchAll", GetarticlesAPI)
	router.POST("/article/searchById", GetarticleAPI)
	router.POST("/article/update", ModarticleAPI)
	router.POST("/article/delete", DelarticleAPI)
	router.Run(":8000")
}
