package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// GetUsers - GET /user
// Get all users
func GetUsers(c *gin.Context) {
	var users []models.User

	models.DB.Preload("CreditCards").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUser - GET /user/:id
// Get a user with UUID
func GetUser(c *gin.Context) {
	var user models.User

	uuid := c.Param("id")

	err := models.DB.Where("id = ?", uuid).Preload("CreditCards").First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with UUID: " + uuid + " not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser - POST /user
// Creates a new user
func CreateUser(c *gin.Context) {
	var newUser models.NewUser
	var user models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DB.Where("email = ?", newUser.Email).First(&user).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with that email already exists"})
		return
	}

	if newUser.Password != newUser.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user = models.User{
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Phone:     newUser.Phone,
		Email:     newUser.Email,
		Password:  hashedPassword,
		Role:      "user",
	}
	result := models.DB.Omit("CreditCards").Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// UpdateUser - PUT /user/:id
// Updates a user
func UpdateUser(c *gin.Context) {
	var user models.User

	uuid := c.Param("id")

	err := models.DB.Where("id = ?", uuid).Preload("CreditCards").First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with UUID: " + uuid + " not found"})
		return
	}

	// TODO: logic for updating a user
}

// DeleteUser - DELETE /user/:id
// Deletes a user
func DeleteUser(c *gin.Context) {
	var user models.User

	uuid := c.Param("id")

	err := models.DB.Where("id = ?", uuid).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with UUID: " + uuid + " not found"})
		return
	}

	models.DB.Delete(&user)

	c.Status(http.StatusNoContent)
}

// LoginUser - POST /user/login
// Returns user when provided credentials are valid
func LoginUser(c *gin.Context) {
	var credentials models.SignIn

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := models.DB.First(&user, "email = ?", strings.ToLower(credentials.Email))
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or Password"})
		return
	}

	if err := utils.VerifyPassword(user.Password, credentials.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Email or Password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
