package main

import (
	"log"
	"github.com/gin-gonic/gin"
)
func main(){
	router:=gin.Default()
	router.GET("/",func(c *gin.Context){
		c.JSON(200,gin.H{
			"message":"Welcome to the Grocery Store API!",
		})
	})

	log.Fatal(router.Run(":8080")) // Run server on port 8080

}