package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaigner/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCampaigns - GET /campaign
// Get all campaigns
func GetCampaigns(c *gin.Context) {
	var campaigns []models.Campaign

	models.DB.Find(&campaigns)

	c.JSON(http.StatusOK, gin.H{"data": campaigns})
}
