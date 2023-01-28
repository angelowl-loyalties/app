package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaigner/models"
	"github.com/gin-gonic/gin"
	"net/http"
	// "io/ioutil"
)

// GetCampaigns - GET /campaign
// @Summary Get all campaigns
// @Description Get all campaigns
// @Tags campaign
// @Produce json
// @Success 200 {array} models.Campaign
// @Router /campaign [get]
func GetCampaigns(c *gin.Context) {
	var campaigns []models.Campaign

	models.DB.Find(&campaigns)

	c.JSON(http.StatusOK, gin.H{"data": campaigns})
}

// GetCampaignById - GET /campaign/:id
// @Summary Get a campaign
// @Description Get a single campaign by its UUID
// @Tags campaign
// @Produce json
// @Success 200 {object} models.Campaign
// @Param campaign_id path string true "Campaign ID"
// @Router /campaign/{campaign_id} [get]
func GetCampaignById(c *gin.Context) {
	var campaign models.Campaign

	id := c.Param("id")

	models.DB.Where("id = ?", id).Find(&campaign)

	c.JSON(http.StatusOK, gin.H{"data": campaign})
}

// CreateCampaign - POST /campaign
// @Summary Create a campaign
// @Description Create a campaign
// @Tags campaign
// @Accept json
// @Produce json
// @Success 200 {object} models.Campaign
// @Param campaign body models.Campaign true "New Campaign"
// @Router /campaign [post]
func CreateCampaign(c *gin.Context) {
	var newCampaign models.Campaign

	if err := c.ShouldBindJSON(&newCampaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Create(&newCampaign)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newCampaign})
}

// UpdateCampaign - PUT /campaign/:id
// @Summary Update a campaign
// @Description Update a campaign
// @Tags campaign
// @Accept json
// @Produce json
// @Success 200 {object} models.Campaign
// @Param campaign body models.Campaign true "New Campaign"
// @Param campaign_id path string true "Campaign ID"
// @Router /campaign/{campaign_id} [put]
func UpdateCampaign(c *gin.Context) {
	var campaign models.Campaign
	var updatedCampaign models.Campaign

	uuid := c.Param("id")

	err := models.DB.Where("id = ?", uuid).First(&campaign).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign with UUID: " + uuid + " not found"})
		return
	}

	if err = c.ShouldBindJSON(&updatedCampaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	campaign.Name = updatedCampaign.Name
	campaign.MinSpend = updatedCampaign.MinSpend
	campaign.Start = updatedCampaign.Start
	campaign.End = updatedCampaign.End
	campaign.RewardProgram = updatedCampaign.RewardProgram
	campaign.RewardAmount = updatedCampaign.RewardAmount
	campaign.MCC = updatedCampaign.MCC
	campaign.Merchant = updatedCampaign.Merchant

	result := models.DB.Save(&campaign)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": campaign})
}

// DeleteCampaign - DELETE /campaign/:id
// @Summary Delete a campaign
// @Description Delete a campaign
// @Tags campaign
// @Produce json
// @Success 204 {object} nil
// @Param campaign_id path string true "Campaign ID"
// @Router /campaign/{campaign_id} [delete]
func DeleteCampaign(c *gin.Context) {
	var campaign models.Campaign

	uuid := c.Param("id")

	err := models.DB.Where("id = ?", uuid).First(&campaign).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign with UUID: " + uuid + " not found"})
		return
	}

	models.DB.Delete(&campaign)

	c.Status(http.StatusNoContent)
}
