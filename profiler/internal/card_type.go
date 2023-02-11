package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCardTypes - GET /card/type
// @Summary Get all card types
// @Description Get all card types
// @Tags card_type
// @Produce json
// @Success 200 {array} models.CardType
// @Router /card/type [get]
func GetCardTypes(c *gin.Context) {
	var cardTypes []models.CardType

	cardTypes, err := models.CardTypeGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cardTypes})
}

// GetCardTypePK - GET /card/type/:cardType
// @Summary Get a card type
// @Description Get a single card type by its PK
// @Tags card_type
// @Produce json
// @Success 200 {object} models.CardType
// @Param card_type path string true "Card Type PK"
// @Router /card/type/{card_type} [get]
func GetCardTypePK(c *gin.Context) {
	reqCardType := c.Param("cardType")

	cardType, err := models.CardTypeGetByType(reqCardType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if cardType == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": cardType})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cardType})
}

// CreateCardType - POST /card/type
// @Summary Create a card type
// @Description Create a card type
// @Tags card_type
// @Accept json
// @Produce json
// @Success 201 {object} models.CardType
// @Param card_type body models.CardType true "New Card Type"
// @Router /card/type [post]
func CreateCardType(c *gin.Context) {
	var cardType models.CardType

	if err := c.ShouldBindJSON(&cardType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCardType, err := models.CardTypeCreate(&cardType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdCardType})
}

// UpdateCardType - PUT /card/type/:cardType
// @Summary Update a card type
// @Description Update a card type
// @Tags card_type
// @Accept json
// @Produce json
// @Success 200 {object} models.CardType
// @Param card_type body models.CardType true "New Card Type"
// @Param card_type_pk path string true "Campaign Type PK"
// @Router /card/type/{card_type_pk} [put]
func UpdateCardType(c *gin.Context) {
	var updatedCardType models.CardType

	reqCardType := c.Param("cardType")
	cardType, err := models.CardTypeGetByType(reqCardType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if cardType == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": cardType})
		return
	}

	if err = c.ShouldBindJSON(&updatedCardType); err != nil || updatedCardType.CardType != reqCardType {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cardType.RewardProgram = updatedCardType.RewardProgram
	cardType.Name = updatedCardType.Name
	cardType.RewardUnit = updatedCardType.RewardUnit

	cardType, err = models.CardTypeSave(cardType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cardType})
}

// DeleteCardType - DELETE /card/type/:cardType
// @Summary Delete a card type
// @Description Delete a card type
// @Tags card_type
// @Produce json
// @Success 204 {object} nil
// @Param card_type_pk path string true "Card Type PK"
// @Router /card/type/{card_type_pk} [delete]
func DeleteCardType(c *gin.Context) {
	reqCardType := c.Param("cardType")

	cardType, err := models.CardTypeGetByType(reqCardType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if cardType == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": cardType})
		return
	}

	_, err = models.CardTypeDelete(cardType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
