package handler

import (
	"net/http"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {

	userWithNewData := &struct {
		User                     database.User
		NewUsername, NewPassword string
	}{}

	if userWithNewData == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	err := database.UpdateUser(&userWithNewData.User, userWithNewData.NewUsername, userWithNewData.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}
}
