package internal

import (
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/models"
	"github.com/cs301-itsa/project-2022-23t2-g1-t7/profiler/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
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

	users, err := models.UserGetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

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
	uuid := c.Param("id")

	user, err := models.UserGetById(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": user})
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
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	temp, err := models.UserGetByEmail(newUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if temp != nil && newUser.Email == temp.Email {
		// check if another user has the same email, if so, error
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with that email already exists"})
		return
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userID, _ := uuid.Parse(newUser.ID)
	user := &models.User{
		ID:        userID,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Phone:     newUser.Phone,
		Email:     newUser.Email,
		Password:  hashedPassword,
		Role:      "user",
	}
	user, err = models.UserCreate(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	var updatedUser models.UserInput

	uuid := c.Param("id")
	user, err := models.UserGetById(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": user})
		return
	}

	if err = c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	temp, err := models.UserGetByEmail(updatedUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
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

	user, err = models.UserSave(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	uuid := c.Param("id")
	user, err := models.UserGetById(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"data": user})
		return
	}

	_, err = models.UserDelete(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
