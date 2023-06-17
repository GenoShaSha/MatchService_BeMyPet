package main

import (
	"machingMicroService/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	tableMatches := "matches" 

	router.POST("/makematch", func(c *gin.Context){
		service.MakeMatch(c, tableMatches)
	})
	router.POST("/getadoptermatches", func(c *gin.Context){
		service.GetAdopterMatches(c, tableMatches)
	})
	router.POST("/getsheltermatches", func(c *gin.Context){
		service.GetShelterMatches(c, tableMatches)})

	router.PUT("/updatematch", func(c *gin.Context){
		service.UpdateMatchStatus(c, tableMatches)
	})
	router.GET("/matches", func(c *gin.Context){
		service.GetMatches(c, tableMatches)
	})

	// Run the router
	router.Run(":8080")

	// kafkaaccess.ConnectAndWriteMessage()
	// kafkaaccess.ConnectAndConsumeMessage()

	//test SonarCube

}
