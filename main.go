package main

import (
	"machingMicroService/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/makematch", service.MakeMatch)
	router.GET("/matches", service.GetMatches)
	router.GET("/update", service.Get)
	// router.GET("/products", GetProducts)
	// router.GET("/products/:productId", GetSingleProduct)
	// router.PUT("/products/:productId", UpdateProduct)
	// router.DELETE("/products/:productId", DeleteProduct)

	// Run the router
	router.Run(":8080")

	// kafkaaccess.ConnectAndWriteMessage()
	// kafkaaccess.ConnectAndConsumeMessage()

}
