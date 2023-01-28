package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/campaigner/models"
	"github.com/gin-gonic/gin"
	"net/http"
	// "io/ioutil"
)

// GetCampaigns - GET /campaign
// Get all campaigns
// @Summary Get all campaigns
// @Schemes
// @Description Get all campaigns
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /campaign [get]
func GetCampaigns(c *gin.Context) {
	var campaigns []models.Campaign
	models.DB.Find(&campaigns)
	c.JSON(http.StatusOK, gin.H{"data": campaigns})
}

func GetCampaignById(c *gin.Context){
	var campaign models.Campaign
	id := c.Param("id")
	models.DB.Where("id = ?",id).Find(&campaign)
	c.JSON(http.StatusOK, gin.H{"data":campaign})
}

func CreateCampaign(c *gin.Context){
	//TODO: Error Checking for mandatory fields
	var campaignBody models.Campaign
	if err := c.BindJSON(&campaignBody); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	err:= models.DB.Create(&campaignBody).Error

	if  (err == nil){
		c.JSON(200, gin.H{"message": "Success"})
		return
	}

	//TODO: Change to more descriptive error 
	//(Not sure what scenario it will end up here)
	c.JSON(500, gin.H{"error":err.Error()})
	return
}