package controllers

import (
	"go-learn-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	loginData := make(map[string]string)
	if err := c.BindJSON(&loginData); err != nil {
		c.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Sorry Server Error"},
		)
		return
	}

	users := models.GetUserLogin(loginData["userId"], loginData["email"])
	c.IndentedJSON(
		http.StatusOK,
		gin.H{"users": users},
	)
}
