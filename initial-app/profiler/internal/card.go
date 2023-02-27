package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

//var validate = validator.New()

// GetCards - GET /card
// @Summary Get all cards
// @Description Get all cards
// @Tags card
// @Produce json
// @Success 200 {array} models.Card
// @Router /card [get]
func GetCards(c *gin.Context) {
	var cards []models.Card

	cards, err := models.CardGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cards})
}

// GetCard - GET /card/:id
// @Summary Get a card
// @Description Get a single card by its UUID
// @Tags card
// @Produce json
// @Success 200 {object} models.Card
// @Param card_id path string true "Card ID"
// @Router /card/{card_id} [get]
func GetCard(c *gin.Context) {
	uuid := c.Param("id")

	card, err := models.CardGetById(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if card == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": card})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": card})
}

// CreateCard - POST /card
// @Summary Create a card
// @Description Create a card
// @Tags card
// @Accept json
// @Produce json
// @Success 201 {object} models.Card
// @Param card body models.Card true "New Card"
// @Router /card [post]
func CreateCard(c *gin.Context) {
	var newCard models.Card

	if err := c.ShouldBindJSON(&newCard); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	temp, err := models.CardGetByPan(newCard.CardPan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if temp != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Card with that PAN already exists"})
		return
	}

	userID := newCard.UserID.String()
	user, err := models.UserGetById(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User with UUID: " + userID})
		return
	}

	reqCardType := newCard.CardTypeCardType
	cardType, err := models.CardTypeGetByType(reqCardType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if cardType != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Card Type: " + reqCardType})
		return
	}

	card, err := models.CardCreate(&newCard)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": card})
}

// DeleteCard - DELETE /card/:id
// @Summary Delete a card
// @Description Delete a card
// @Tags card
// @Produce json
// @Success 204 {object} nil
// @Param card_id path string true "Card ID"
// @Router /card/{card_id} [delete]
func DeleteCard(c *gin.Context) {
	uuid := c.Param("id")

	card, err := models.CardGetById(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if card == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": card})
		return
	}

	_, err = models.CardDelete(card)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}