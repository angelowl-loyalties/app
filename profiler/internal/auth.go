package internal

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/utils"
)

// LoginUser - POST /user/login
// @Summary User login
// @Description Returns user when provided credentials are valid
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.AuthResponse
// @Param credentials body models.SignIn true "Credentials"
// @Router /auth/login [post]
func LoginUser(c *gin.Context) {
	var credentials models.SignIn

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	user, err := models.UserGetByEmail(strings.ToLower(credentials.Email))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or Password"})
		return
	}

	if err := utils.VerifyPassword(user.Password, credentials.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or Password"})
		return
	}

	jwtToken := utils.CreateJWT(user)

	signedJWT, err := jwtToken.SignedString(utils.KMSConfig.WithContext(context.Background()))
	if err != nil {
		log.Println("failed to sign JWT: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := models.AuthResponse{
		JWT:    signedJWT,
		UserID: user.ID.String(),
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}
