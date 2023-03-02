package main

import (
	"fmt"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaignex/internal"
	"log"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaignex/config"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaignex/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaignex/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/cs301-itsa/project-2022-23t2-g1-t7/campaignex/docs"
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

	// setup DB connection
	dbConnString := c.DBConnString
	models.ConnectDB(dbConnString)

	// setup etcd connection
	etcdEndpoints := c.EtcdEndpoints
	internal.InitClient(etcdEndpoints)

	// setup routes
	router = gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.InitialiseRoutes(router)

	port := c.Port
	_ = router.Run(":" + port)
}
