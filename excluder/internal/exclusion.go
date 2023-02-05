package internal

import (
	"net/http"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/excluder/models"
	"github.com/gin-gonic/gin"
)

// GetExclusions - GET /exclusion
// @Summary Get all exclusions
// @Description Get all exclusions
// @Tags exclusion
// @Produce json
// @Success 200 {array} models.Exclusion
// @Router /exclusion [get]
func GetExclusions(c *gin.Context) {
	exclusions, err := models.ExclusionGetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": exclusions})
}

// GetExclusionById - GET /exclusion/:id
// @Summary Get an exclusion
// @Description Get a single exclusion by its UUID
// @Tags exclusion
// @Produce json
// @Success 200 {object} models.Exclusion
// @Param exclusion_id path string true "Exclusion ID"
// @Router /exclusion/{exclusion_id} [get]
func GetExclusionById(c *gin.Context) {
	uuid := c.Param("id")

	exclusion, err := models.ExclusionGetById(uuid)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exclusion == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exclusion})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": exclusion})
}

// CreateExclusion - POST /exclusion
// @Summary Create an exclusion
// @Description Create an exclusion
// @Tags exclusion
// @Accept json
// @Produce json
// @Success 200 {object} models.Exclusion
// @Param exclusion body models.Exclusion true "New Exclusion"
// @Router /exclusion [post]
func CreateExclusion(c *gin.Context) {
	var newExclusion models.Exclusion

	if err := c.ShouldBindJSON(&newExclusion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exclusion, err := models.ExclusionCreate(&newExclusion)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": exclusion})
}

// UpdateExclusion - PUT /exclusion/:id
// @Summary Update a exclusion
// @Description Update a exclusion
// @Tags exclusion
// @Accept json
// @Produce json
// @Success 200 {object} models.Exclusion
// @Param exclusion body models.Exclusion true "New Exclusion"
// @Param exclusion_id path string true "Exclusion ID"
// @Router /exclusion/{exclusion_id} [put]
func UpdateExclusion(c *gin.Context) {
	var updatedExclusion models.Exclusion

	uuid := c.Param("id")

	exclusion, err := models.ExclusionGetById(uuid)

	// Error in DB during ExclusionGetById
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Not found
	if exclusion == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exclusion})
		return
	}

	if err = c.ShouldBindJSON(&updatedExclusion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exclusion.MCC = updatedExclusion.MCC

	exclusion, err = models.ExclusionSave(exclusion)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": exclusion})
}

// DeleteExclusion - DELETE /exclusion/:id
// @Summary Delete a exclusion
// @Description Delete a exclusion
// @Tags exclusion
// @Produce json
// @Success 204 {object} nil
// @Param exclusion_id path string true "Exclusion ID"
// @Router /exclusion/{exclusion_id} [delete]
func DeleteExclusion(c *gin.Context) {
	uuid := c.Param("id")

	exclusion, err := models.ExclusionGetById(uuid)

	// Error in DB during ExclusionGetById
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Not found
	if exclusion == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": exclusion})
		return
	}

	_, err = models.ExclusionDelete(exclusion)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
