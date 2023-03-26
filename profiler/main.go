package main

import (
	"fmt"
	"log"
	"time"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/config"
	_ "github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/docs"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/routes"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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

	utils.InitKMS(c.AWSAccessKey, c.AWSSecretKey, c.JWTKMSKey)

	router = gin.Default()
	router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "PUT", "DELETE", "POST"},
        AllowCredentials: true,
        MaxAge: 12 * time.Hour,
    }))

	// docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routes.InitialiseRoutes(router)

	port := c.Port
	_ = router.Run(":" + port)
}
