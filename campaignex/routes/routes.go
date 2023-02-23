package routes

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaignex/internal"
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

	campaign := router.Group("/campaign")
	{
		campaign.GET("/", internal.GetCampaigns)
		campaign.GET("/:id", internal.GetCampaignById)
		campaign.POST("/", internal.CreateCampaign)
		campaign.PUT("/:id", internal.UpdateCampaign)
		campaign.DELETE("/:id", internal.DeleteCampaign)
	}

    exclusion := router.Group("/exclusion")
    {
        exclusion.GET("/", internal.GetExclusions)
        exclusion.GET("/:id", internal.GetExclusionById)
        exclusion.POST("/", internal.CreateExclusion)
        exclusion.PUT("/:id", internal.UpdateExclusion)
        exclusion.DELETE("/:id", internal.DeleteExclusion)
    }
}
