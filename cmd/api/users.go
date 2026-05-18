package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} database.User
// @Failure 500 {object} gin.H{"error": "Something went wrong, Failed to retrieve users"}
// @Router /users [get]


func (app *application) getAllUsers(c *gin.Context) {
	users, err := app.models.Users.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong, Failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, users)
}