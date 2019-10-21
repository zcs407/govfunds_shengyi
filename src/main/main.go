package main

import (
	"github.com/gin-gonic/gin"
	"govfunds_shengyi/src/common"
	"govfunds_shengyi/src/controller"
	"log"
)

func main() {
	common.InitES()
	common.InitDB()
	router := gin.Default()
	router.Use(gin.Recovery())
	r1 := router.Group("/govfunds")
	{
		r1.POST("/search", controller.Search)
		r1.PUT("/dataSync2es", controller.DataSync2es)
	}
	err := router.Run(":8080")
	if err != nil {
		log.Println("zf_info_search service running err:", err)
	}
}
