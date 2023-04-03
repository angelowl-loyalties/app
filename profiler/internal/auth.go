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

// LoginUser - POST /auth/login
// @Summary User login
// @Description Returns user when provided credentials are valid
// @Tags auth
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
		IsNew:  user.IsNew,
	}

	c.JSON(http.StatusOK, gin.H{"data": response})
}

// ChangeDefaultPassword - POST /auth/password
// @Summary Change default password
// @Description Endpoint allows a user to change their default password
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Param credentials body models.ChangeDefaultPassword true "Credentials"
// @Router /auth/password [post]
func ChangeDefaultPassword(c *gin.Context) {
	// validate input struct
	var newCredentials models.ChangeDefaultPassword
	if err := c.ShouldBindJSON(&newCredentials); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// check that the user exists
	user, err := models.UserGetByEmail(strings.ToLower(newCredentials.Email))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or Password"})
		return
	}

	// validate the user is correct by checking the old password is the same
	if err := utils.VerifyPassword(user.Password, newCredentials.OldPassword); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or Password"})
		return
	}

	// check that the user is new
	if !user.IsNew {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User is not new"})
		return
	}

	hashedPassword, err := utils.HashPassword(newCredentials.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = hashedPassword
	user.IsNew = false

	_, err = models.UserSave(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "OK"})
}
