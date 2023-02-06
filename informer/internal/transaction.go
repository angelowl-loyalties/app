package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetRewards - GET /reward
// @Summary Get all rewards
// @Description Get all rewards
// @Tags reward
// @Produce json
// @Success 200 {array} models.Reward
// @Router /reward [get]
func GetRewards(c *gin.Context) {
	var rewards []models.Reward

	rewards, _ = models.RewardGetAll()

	c.JSON(http.StatusOK, gin.H{"data": rewards})
}
