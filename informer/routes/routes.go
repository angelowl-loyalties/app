package routes

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/internal"
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

	transaction := router.Group("/reward")
	{
		transaction.GET("/", internal.GetRewardsCount)
		transaction.GET("/:cardId", internal.GetRewardsByCardID)
		transaction.GET("/total/:cardId", internal.GetTotalRewardsByCardID)
	}
}
