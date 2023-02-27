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
	// this function should be deprecated and removed in the future
	var rewards []models.Reward

	rewards, err := models.RewardGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rewards})
}

// GetRewardsByCardID - GET /reward/:cardId
// @Summary Get all rewards for a specified card
// @Description Get all rewards for a particular card's UUID
// @Tags reward
// @Produce json
// @Success 200 {array} models.Reward
// @Param cardId path string true "Card ID"
// @Router /reward/{cardId} [get]
func GetRewardsByCardID(c *gin.Context) {
	var rewards []models.Reward
	cardID := c.Param("cardId")

	rewards, err := models.RewardGetByCardID(cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rewards})
}
