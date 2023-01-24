package routes

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/internal"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func InitialiseRoutes(router *gin.Engine) {
	router.GET("/health", Health)

	//user := router.Group("/user")
	//{
	//	user.GET("/", internal.GetUsers)
	//	user.GET("/:id", internal.GetUser)
	//	user.POST("/", internal.CreateUser)
	//	user.DELETE("/:id", internal.DeleteUser)
	//}

	cardType := router.Group("/card/type")
	{
		cardType.GET("/", internal.GetCardTypes)
		cardType.GET("/:cardType", internal.GetCardTypePK)
		cardType.POST("/", internal.CreateCardType)
		cardType.PUT("/:cardType", internal.UpdateCardType)
		cardType.DELETE("/:cardType", internal.DeleteCardType)
	}

	//card := router.Group("/card")
	//{
	//	card.GET("/")
	//}
}
