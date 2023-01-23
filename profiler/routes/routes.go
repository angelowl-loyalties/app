package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func InitialiseRoutes(router *gin.Engine) {
	router.GET("/health", Health)

	//user := router.Group("/user")
	//
	//card := router.Group("/card")
	//
	//cardType := router.Group("/card/type")
}
