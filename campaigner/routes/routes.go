package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaigner/internal"
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

	campaign := router.Group("/campaign")
	campaign.POST("/",internal.CreateCampaign)
	campaign.GET("/",internal.GetCampaigns)
	campaign.GET("/:id",internal.GetCampaignById)

	// add your other routes accordingly
}
