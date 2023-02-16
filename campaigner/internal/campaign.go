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

	err = etcdPutCampaign(campaign) // put into etcd
	if err != nil {
		// roll back persist to DB if put to etcd fails
		_, _ = models.CampaignDelete(campaign)
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

	originalCampaign := campaign
	campaign.Name = updatedCampaign.Name
	campaign.MinSpend = updatedCampaign.MinSpend
	campaign.Start = updatedCampaign.Start
	campaign.End = updatedCampaign.End
	campaign.RewardProgram = updatedCampaign.RewardProgram
	campaign.RewardAmount = updatedCampaign.RewardAmount
	campaign.MCC = updatedCampaign.MCC
	campaign.Merchant = updatedCampaign.Merchant

	// put into etcd, nothing to roll back if failure
	err = etcdPutCampaign(campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	campaign, err = models.CampaignSave(campaign)
	if err != nil {
		// since persist to DB fails, we roll back update to etcd
		_ = etcdPutCampaign(originalCampaign)
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

	_, err = etcdDeleteCampaign(campaign.ID.String()) // delete from etcd
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_, err = models.CampaignDelete(campaign)
	if err != nil {
		// since deletion fails, we need to restore etcd copy of campaign
		_ = etcdPutCampaign(campaign)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
