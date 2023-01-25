package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/models"
	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"
	"net/http"
)

//var validate = validator.New()

// GetCards - GET /card
// Get all cards
func GetCards(c *gin.Context) {
	var cards []models.Card

	models.DB.Find(&cards)

	c.JSON(http.StatusOK, gin.H{"data": cards})
}

// GetCard - GET /card/:id
// Get a card given UUID
func GetCard(c *gin.Context) {
	var card models.Card

	uuid := c.Param("id")

	err := models.DB.Where("id = ?", uuid).First(&card).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card with UUID: " + uuid + " not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": card})
}

// CreateCard - POST /card
// Creates a new card
func CreateCard(c *gin.Context) {
	var newCard models.Card
	var user models.User
	var cardType models.CardType

	if err := c.BindJSON(&newCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var temp models.Card
	err := models.DB.Where("card_pan = ?", newCard.CardPan).First(&temp).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Card with that PAN already exists"})
		return
	}

	userID := newCard.UserID.String()
	err = models.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User with UUID: " + userID})
		return
	}

	reqCardType := newCard.CardTypeCardType
	err = models.DB.Where("card_type = ?", reqCardType).First(&cardType).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card Type: " + reqCardType + " not found"})
		return
	}

	result := models.DB.Create(&newCard)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": newCard})
}

// DeleteCard - DELETE /card/:id
// Deletes a card
func DeleteCard(c *gin.Context) {
	var card models.Card

	uuid := c.Param("id")

	err := models.DB.Where("id = ?", uuid).First(&card).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card with UUID: " + uuid + " not found"})
		return
	}

	models.DB.Delete(&card)

	c.Status(http.StatusNoContent)
}
