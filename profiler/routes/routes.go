package routes

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health godoc
// @Summary health
// @Description health check
// @Tag health
// @Produce json
// @Success 200 {string} OK
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func InitialiseRoutes(router *gin.Engine) {
	router.GET("/health", Health)

	auth := router.Group("/auth")
	{
		auth.POST("/login", internal.LoginUser)
	}

	user := router.Group("/user")
	{
		user.GET("/", internal.GetUsers)
		user.GET("/:id", internal.GetUser)
		user.POST("/", internal.CreateUser)
		user.PUT("/:id", internal.UpdateUser)
		user.DELETE("/:id", internal.DeleteUser)
	}

	cardType := router.Group("/card/type")
	{
		cardType.GET("/", internal.GetCardTypes)
		cardType.GET("/:cardType", internal.GetCardTypePK)
		cardType.POST("/", internal.CreateCardType)
		cardType.PUT("/:cardType", internal.UpdateCardType)
		cardType.DELETE("/:cardType", internal.DeleteCardType)
	}

	card := router.Group("/card")
	{
		card.GET("/", internal.GetCards)
		card.GET("/:id", internal.GetCard)
		card.POST("/", internal.CreateCard)
		card.DELETE("/:id", internal.DeleteCard)
	}
}
