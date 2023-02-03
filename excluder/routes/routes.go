package routes

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/excluder/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health godoc
// @Summary health
// @Description health check
// @Produce json
// @Success 200 {string} OK
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func InitialiseRoutes(router *gin.Engine) {
	router.GET("/health", Health)

	exclusion := router.Group("/exclusion")
	{
		exclusion.GET("/", internal.GetExclusions)
	// 	campaign.GET("/:id", internal.GetExclusionById)
	// 	campaign.POST("/", internal.CreateExclusion)
	// 	campaign.PUT("/:id", internal.UpdateExclusion)
	// 	campaign.DELETE("/:id", internal.DeleteExclusion)
	}
	// add your other routes accordingly
}
