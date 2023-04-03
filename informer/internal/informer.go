package internal

import (
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/informer/models"
	"github.com/gin-gonic/gin"
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

	c.JSON(http.StatusOK, gin.H{"total_rewards": len(rewards)})
	//c.JSON(http.StatusOK, gin.H{"data": rewards, "total_rewards": len(rewards)})
}

// GetRewardsByCardID - GET /reward/:cardId
// @Summary Get all rewards for a specified card
// @Description Get all rewards for a particular card's UUID
// @Tags reward
// @Produce json
// @Success 200 {array} models.Reward
// @Param cardId path string true "Card ID"
// @Param page_size query string false "Page Size"
// @Param page_no query string false "Page Number"
// @Router /reward/{cardId} [get]
func GetRewardsByCardID(c *gin.Context) {
	var rewards []models.Reward
	cardID := c.Param("cardId")

	rewards, err := models.RewardGetByCardID(cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rewardsCount := len(rewards)
	if rewardsCount == 0 {
		// zero rewards
		c.JSON(http.StatusOK, gin.H{"page_no": 1, "total_rewards": rewardsCount, "data": rewards})
		return
	}
	pageSizeReq := c.DefaultQuery("page_size", "20")
	pageNumReq := c.DefaultQuery("page_no", "1")
	pageSize, err := strconv.Atoi(pageSizeReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pageNum, err := strconv.Atoi(pageNumReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// requested page size is more than rewards count, set page size to rewards count, only 1 page in this case
	if pageSize > rewardsCount {
		pageSize = rewardsCount
	}

	// calculate number of pages and check that requested page number is valid
	totalPages := int(math.Ceil(float64(rewardsCount) / float64(pageSize)))
	log.Println("total pages: ", totalPages)
	if pageNum > totalPages || pageNum < 1 {
		// non-zero rewards
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page number"})
		return
	}

	var rewardsSlice []models.Reward
	lowerIndex := pageSize * (pageNum - 1) // multiples of pageSize, starting from 0
	// slice from rewards with generated index note that [inclusive:exclusive]
	// handle last page, where number of rewards are likely less than page size
	if pageNum == totalPages {
		rewardsSlice = rewards[lowerIndex:]
	} else {
		upperIndex := lowerIndex + pageSize
		rewardsSlice = rewards[lowerIndex:upperIndex]
	}

	c.JSON(http.StatusOK, gin.H{"page_no": pageNum, "total_rewards": rewardsCount, "data": rewardsSlice})
}

// GetTotalRewardsByCardID - GET /reward/total/:cardId
// @Summary Get total rewards for a specified card
// @Description Get total rewards given a particular card's UUID
// @Tags reward
// @Produce json
// @Success 200 {array} number
// @Param cardId path string true "Card ID"
// @Router /reward/total/{cardId} [get]
func GetTotalRewardsByCardID(c *gin.Context) {
	var rewards []models.Reward
	cardID := c.Param("cardId")

	rewards, err := models.RewardGetByCardID(cardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalRewards := 0.0
	for _, reward := range rewards {
		totalRewards += reward.RewardAmount
	}

	c.JSON(http.StatusOK, gin.H{"data": totalRewards})
}

// GetRewardsForToday - GET /reward/today
// @Summary Get total rewards created today, with reward_amount > 0
// @Description This endpoint is for email aggregation purposes to notify users of the rewards they earned today
// @Tags reward
// @Produce json
// @Success 200 {array} number
// @Router /reward/today [get]
func GetRewardsForToday(c *gin.Context) {
	var rewards []models.Reward

	rewards, err := models.RewardGetTodays()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": rewards})
}
