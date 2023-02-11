package internal

import (
	"net/http"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaigner/models"
	"github.com/gin-gonic/gin"
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
	campaigns, err := models.CampaignGetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	uuid := c.Param("id")

	campaign, err := models.CampaignGetById(uuid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if campaign == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": campaign})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": campaign})
}

// CreateCampaign - POST /campaign
// @Summary Create a campaign
// @Description Create a campaign
// @Tags campaign
// @Accept json
// @Produce json
// @Success 201 {object} models.Campaign
// @Param campaign body models.Campaign true "New Campaign"
// @Router /campaign [post]
func CreateCampaign(c *gin.Context) {
	var newCampaign models.Campaign

	if err := c.ShouldBindJSON(&newCampaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	campaign, err := models.CampaignCreate(&newCampaign)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": campaign})
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
	var updatedCampaign models.Campaign

	uuid := c.Param("id")

	campaign, err := models.CampaignGetById(uuid)

	// Error in DB during CampaignGetById
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Not found
	if campaign == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": campaign})
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

	campaign, err = models.CampaignSave(campaign)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	uuid := c.Param("id")

	campaign, err := models.CampaignGetById(uuid)

	// Error in DB during CampaignGetById
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Not found
	if campaign == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": campaign})
		return
	}

	_, err = models.CampaignDelete(campaign)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
