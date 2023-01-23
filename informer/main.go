package main

import (
	"fmt"
	"log"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/config"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/routes"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	fmt.Println("Server starting...")

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	dbConnString := c.DBConnString
	dbKeyspace := c.DBKeyspace
	dbUser := c.DBUser
	dbPass := c.DBPass
	models.ConnectDB(dbConnString, dbKeyspace, dbUser, dbPass)

	router = gin.Default()
	routes.InitialiseRoutes(router)
	
	port := c.Port
	router.Run(":" + port)
}
