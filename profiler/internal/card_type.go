package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCardTypes - GET /card/type
// Get all card types
func GetCardTypes(c *gin.Context) {
	var cardTypes []models.CardType

	models.DB.Preload("Cards").Find(&cardTypes)

	c.JSON(http.StatusOK, gin.H{"data": cardTypes})
}

// GetCardTypePK - GET /card/type/:cardType
// Get a card type with PK
func GetCardTypePK(c *gin.Context) {
	var cardType models.CardType

	reqCardType := c.Param("cardType")

	err := models.DB.Where("card_type = ?", reqCardType).Preload("Cards").First(&cardType).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card Type: " + reqCardType + " not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cardType})
}

// CreateCardType - POST /card/type
// Create a new card type
func CreateCardType(c *gin.Context) {
	var cardType models.CardType

	if err := c.ShouldBindJSON(&cardType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Create(&cardType)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": cardType})
}

// UpdateCardType - PUT /card/type/:cardType
// Update a card type
func UpdateCardType(c *gin.Context) {
	var cardType models.CardType
	var updatedCardType models.CardType

	reqCardType := c.Param("cardType")

	err := models.DB.Where("card_type = ?", reqCardType).Preload("Cards").First(&cardType).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card Type: " + reqCardType + " not found"})
		return
	}

	if err = c.ShouldBindJSON(&updatedCardType); err != nil || updatedCardType.CardType != reqCardType {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cardType.CardProgram = updatedCardType.CardProgram
	cardType.Name = updatedCardType.Name
	cardType.RewardUnit = updatedCardType.RewardUnit

	result := models.DB.Save(&cardType)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cardType})
}

// DeleteCardType - DELETE /card/type/:cardType
// Deletes a card type
func DeleteCardType(c *gin.Context) {
	var cardType models.CardType

	reqCardType := c.Param("cardType")

	err := models.DB.Where("card_type = ?", reqCardType).First(&cardType).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card Type: " + reqCardType + " not found"})
		return
	}

	models.DB.Delete(&cardType)

	c.Status(http.StatusNoContent)
}
