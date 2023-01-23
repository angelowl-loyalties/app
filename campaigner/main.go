package main

import (
	"fmt"
	"log"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaigner/config"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaigner/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaigner/routes"
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
	models.ConnectDB(dbConnString)

	port := c.Port

	router = gin.Default()
	routes.InitialiseRoutes(router)
	router.Run(":" + port)
}
