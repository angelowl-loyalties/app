package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// GetUsers - GET /user
// @Summary Get all users
// @Description Get all users
// @Tags user
// @Produce json
// @Success 200 {array} models.User
// @Router /user [get]
func GetUsers(c *gin.Context) {
	var users []models.User

	models.DB.Preload("CreditCards").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUser - GET /user/:id
// @Summary Get a user
// @Description Get a single user by its UUID
// @Tags user
// @Produce json
// @Success 200 {object} models.User
// @Param user_id path string true "User ID"
// @Router /user/{user_id} [get]
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
// @Summary Create a user
// @Description Create a user
// @Tags user
// @Accept json
// @Produce json
// @Success 201 {object} models.User
// @Param user body models.UserInput true "New User"
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var newUser models.UserInput
	var user models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var temp models.User
	err := models.DB.Where("email = ?", newUser.Email).First(&temp).Error
	if newUser.Email == temp.Email {
		// check if another user has the same email, if so, error
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with that email already exists"})
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
// @Summary Update a user
// @Description Update a user
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Param user body models.UserInput true "Updated User"
// @Param user_id path string true "User ID"
// @Router /user/{user_id} [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	var updatedUser models.UserInput

	uuid := c.Param("id")

	err := models.DB.Where("id = ?", uuid).Preload("CreditCards").First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with UUID: " + uuid + " not found"})
		return
	}

	if err = c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var temp models.User
	err = models.DB.Where("email = ?", updatedUser.Email).First(&temp).Error
	if updatedUser.Email != user.Email && updatedUser.Email == temp.Email {
		// if provided email is different from current email, user wants to change their email
		// check if another user has the same email, if so, error
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with that email already exists"})
		return
	}

	hashedPassword, err := utils.HashPassword(updatedUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.FirstName = updatedUser.FirstName
	user.LastName = updatedUser.LastName
	user.Email = updatedUser.Email
	user.Password = hashedPassword

	result := models.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DeleteUser - DELETE /user/:id
// @Summary Delete a user
// @Description Delete a user
// @Tags user
// @Produce json
// @Success 204 {object} nil
// @Param user_id path string true "User ID"
// @Router /user/{user_id} [delete]
func DeleteUser(c *gin.Context) {
	var user models.User

	uuid := c.Param("id")

	err := models.DB.Where("id = ?", uuid).First(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User with UUID: " + uuid + " not found"})
		return
	}

	result := models.DB.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// LoginUser - POST /user/login
// @Summary User login
// @Description Returns user when provided credentials are valid
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Param credentials body models.SignIn true "Credentials"
// @Router /user/login [post]
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
