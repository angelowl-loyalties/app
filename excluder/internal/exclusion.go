package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/excluder/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetExclusions - GET /exclusion
// Get all exclusions
func GetExclusions(c *gin.Context) {
	var exclusions []models.Exclusion

	models.DB.Find(&exclusions)

	c.JSON(http.StatusOK, gin.H{"data": exclusions})
}
