package main

import (
	"fmt"
	"log"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/excluder/config"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/excluder/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/excluder/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/cs301-itsa/project-2022-23t2-g1-t7/excluder/docs"
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
	models.ConnectDB(dbConnString)

	port := c.Port

	router = gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.InitialiseRoutes(router)

	router.Run(":" + port)
}
