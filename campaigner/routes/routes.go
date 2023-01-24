package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingExample godoc
// @Summary healthcheck example
// @Schemes
// @Description health check
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func InitialiseRoutes(router *gin.Engine) {
	router.GET("/health", Health)

	//campaign := router.Group("/campaign")
	// add your other routes accordingly
}
