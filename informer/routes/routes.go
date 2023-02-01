package routes

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func InitialiseRoutes(router *gin.Engine) {
	router.GET("/health", Health)

	transaction := router.Group("/transaction")
	{
		transaction.GET("/", internal.GetAllTransactions)
	}
}
