package main

import (
	"fmt"
	"log"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/config"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	dbTable := c.DBTable
	// dbUser := c.DBUser
	// dbPass := c.DBPass
	models.InitDB(dbConnString, dbKeyspace, dbTable)
	models.ConnectDB(dbConnString, dbKeyspace)

	router = gin.Default()

	// docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.InitialiseRoutes(router)

	port := c.Port
	router.Run(":" + port)
}