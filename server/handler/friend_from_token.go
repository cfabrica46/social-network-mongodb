package handler

import (
	"net/http"

	"github.com/cfabrica46/social-network-mongodb/server/database"
	"github.com/gin-gonic/gin"
)

func GetFriendsFromUser(c *gin.Context) {

	var err error

	user := c.MustGet("user-data").(*database.User)
	if user == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"ErrMessage": "Internal Error",
		})
		return
	}

	friends, err := database.GetFriendsFromUser(user.Friends)
	if err != nil {
		return
	}

	return
}
