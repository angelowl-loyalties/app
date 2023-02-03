package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/excluder/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCampaigns - GET /exclusion
// @Summary Get all exclusion
// @Description Get all exclusions
// @Tags exclusion
// @Produce json
// @Success 200 {array} models.Exclusion
// @Router /exclusion [get]
func GetExclusions(c *gin.Context) {
	var exclusions []models.Exclusion

	models.DB.Find(&exclusions)

	c.JSON(http.StatusOK, gin.H{"data": exclusions})
}
